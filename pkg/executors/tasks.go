package executors

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/utils"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func executeDockerCMD(taskName, command, defaultHive, dockerHive, image string, mounts, inputs []string, timeout time.Duration, displayStdout, dynamic, enablelogs bool) error {
	if defaultHive == "" {
		defaultHive = "hive"
	}
	if dockerHive == "" {
		dockerHive = "/hive"
	}
	iteration := ""
	dockerName := ""
	if dynamic {
		dynamicData := strings.Split(taskName, ":")
		taskName = dynamicData[0]
		iteration = dynamicData[1]
		dockerName = utils.SanitizeTaskName(taskName) + "-" + iteration
	}
	if timeout == 0 {
		timeout = 24 * time.Hour
	}
	sanitizedTaskName := utils.SanitizeTaskName(taskName)
	if !dynamic {
		dockerName = sanitizedTaskName
	}
	hostHiveTaskDir, err := utils.GetHostHiveTaskDirectory(taskName, defaultHive)
	if dynamic {
		hostHiveTaskDir = filepath.Join(hostHiveTaskDir, iteration)
		err = utils.EnsurePathExists(hostHiveTaskDir)
	}
	if err != nil {
		return err
	}
	hostHiveTaskInputDir, hostHiveTaskOutputDir := utils.GetInputOutput(hostHiveTaskDir)
	dockerHiveTaskDir := filepath.Join(dockerHive, sanitizedTaskName)
	var stdoutFile, stderrfile string
	if enablelogs {
		stdoutFile, stderrfile = utils.GeterrorLogPath(hostHiveTaskDir)
	}
	// It will check only length is > 0
	if len(inputs) > 0 {
		copyerror := utils.CopyInputFiles(inputs, hostHiveTaskInputDir)
		if copyerror != nil {
			return copyerror
		}
	}

	err = utils.CopyMountFiles(mounts, hostHiveTaskInputDir, defaultHive)
	if err != nil {
		return err
	}

	command = utils.ReplaceDockerplaceholders(command, dockerHive, dockerHiveTaskDir, hostHiveTaskDir, hostHiveTaskOutputDir, hostHiveTaskInputDir)
	command = utils.ReplaceTaskPlaceholders(command, dockerHive, "static")

	// Create a context with the specified timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create docker client: %v", err)
	}

	if image != "" {
		err = utils.PullDockerImage(image)
		if err != nil {
			return err
		}
	}

	// Docker Command
	cmd, err := utils.InspectImageEntrypoint(image, command)
	if err != nil {
		return fmt.Errorf("error inspecting Docker image entrypoint: %v", err)
	}

	resultVolume := mount.Mount{
		Type:   mount.TypeBind,
		Source: hostHiveTaskDir,
		Target: dockerHive,
	}
	// Check if docker container with same name exist and delete if exist
	// **Change:** Only delete the container if it's not running
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return fmt.Errorf("failed to list containers: %v", err)
	}
	for i := range containers {
		containerObj := &containers[i]
		if containerObj.Names[0] == "/"+dockerName {
			if containerObj.State != "running" {
				err := cli.ContainerRemove(ctx, containerObj.ID, types.ContainerRemoveOptions{Force: true})
				if err != nil {
					return fmt.Errorf("failed to remove container '%s': %v", dockerName, err)
				}
			}
		}
	}

	// Create a container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   cmd,
	}, &container.HostConfig{
		Mounts:     []mount.Mount{resultVolume},
		Privileged: true,
	}, nil, nil, dockerName)
	if err != nil {
		return fmt.Errorf("failed to create container for task '%s': %v", sanitizedTaskName, err)
	}

	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container for task '%s': %v", taskName, err)
	}
	// Wait for the container to finish
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("error while waiting for container for task '%s': %v", taskName, err)
		}
	case <-statusCh:
	}

	// Get the logs from the container
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return fmt.Errorf("failed to get logs for task '%s': %v", taskName, err)
	}
	defer out.Close()
	// Parse the logs to separate stdout and stderr
	stdoutLogs, stderrLogs, err := utils.ParseDockerLogs(out)
	if err != nil {
		return fmt.Errorf("failed to parse logs for task '%s': %v", taskName, err)
	}

	// Optionally display stdout
	if displayStdout {
		if len(stdoutLogs) > 0 {
			fmt.Print(visuals.PrintRandomColor(string(stdoutLogs), 32)) // green
		}
		if len(stderrLogs) > 0 {
			fmt.Print(visuals.PrintRandomColor(string(stderrLogs), 31)) // red
		}

	}
	if enablelogs {
		utils.WriteLogs(stdoutFile, stderrfile, stdoutLogs, stderrLogs)
	}

	// Remove the container once it's finished
	if err := cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true}); err != nil {
		return fmt.Errorf("failed to remove container for task '%s': %v", taskName, err)
	}

	return nil
}

func ExecHelper(configuration *config.Config) {
	// Tasks are false by default
	taskStatus, whitelist := getTaskStatusandWhitelist(configuration)
	// Create a wait group for parallel tasks
	var wg sync.WaitGroup
	// Pull Docker images
	for i := range configuration.Tasks {
		task := &configuration.Tasks[i]
		if task.Image != "" {
			err := utils.PullDockerImage(task.Image)
			if err != nil {
				visuals.PrintState("FATAL", task.Name, err.Error())
				return
			}
		}
	}

	// Function to execute a task
	executeTaskFunc := func(taskName string, taskData interface{}) {
		visuals.PrintState("START", taskName, "")
		err := executeTask(configuration, taskName, taskData)
		if err == nil {
			visuals.PrintState("SUCCESS", taskName, "")
			taskStatus[taskName] = true
		}
	}

	for i := range configuration.Tasks {
		taskData := &configuration.Tasks[i]
		checkRequirements(taskData, whitelist)

		for _, requiredTask := range taskData.Required {
			for !taskStatus[requiredTask] {

				time.Sleep(time.Second)
			}
		}

		if taskData.Parallel {
			// Execute parallel tasks without waiting
			wg.Add(1)
			go func(name string, data interface{}) {
				executeTaskFunc(name, data)
				wg.Done()
			}(taskData.Name, GetTaskDataWithName(taskData.Name, configuration))
		} else {
			// Function to execute tasks with dependencies
			executeWithDependencies := func() {
				// Check if all required tasks are completed
				allRequirementsMet := true
				for _, requiredTask := range taskData.Required {
					if !taskStatus[requiredTask] {
						allRequirementsMet = false
						break
					}
				}

				// Execute the current task if all requirements are met
				if allRequirementsMet {
					executeTaskFunc(taskData.Name, GetTaskDataWithName(taskData.Name, configuration))
				}
			}

			// Execute tasks with dependencies sequentially
			executeWithDependencies()

			// Wait for 1 second before starting the next task
			time.Sleep(time.Second)
		}
	}

	// Wait for all parallel tasks to finish
	wg.Wait()
}

func executeSingleTask(taskName string, commands []string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, dockerimage, dockerHive string, mounts, inputs []string) error {
	visuals.PrintState("Task-Info", taskName, "Task is Static")
	visuals.PrintState("Static-Task: "+taskName, taskName, "Executing Task")

	err := executeDockerCMD(taskName, strings.Join(commands, " && "), wfc.DefaultHive, dockerHive, dockerimage, mounts, inputs, timeout, silent, false, wfc.EnableLogs)

	if err != nil {
		if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "waiting for container") {
			visuals.PrintState("TIMEOUT", taskName, err.Error())
		} else {
			visuals.PrintState("ERROR", taskName, err.Error())
		}
		if stop {

			visuals.PrintState("FATAL", taskName, err.Error())
			return err
		}
		return err
	}
	return nil
}

func executeTask(configuration *config.Config, taskName string, taskData interface{}) error {
	commands := utils.GetInterfaceVal(taskData, "cmds").([]string)
	timeout := utils.GetTimeout(taskData)
	silent := utils.GetInterfaceVal(taskData, "silent").(bool)
	silent = !silent
	stop := utils.GetInterfaceVal(taskData, "stop").(bool)
	wfc := &configuration.WorkflowConfig
	dockerImage := utils.GetInterfaceVal(taskData, "image").(string)
	dockerhive := utils.GetInterfaceVal(taskData, "dockerHive").(string)
	mounts := utils.GetInterfaceVal(taskData, "inputMounts").([]string)
	inputs := utils.GetInterfaceVal(taskData, "inputs").([]string)
	// default docker image from workflow config
	if dockerImage == "" {
		dockerImage = configuration.WorkflowConfig.DefaultDockerimage
	}
	taskType := utils.GetInterfaceVal(taskData, "type").(string)
	if strings.EqualFold(taskType, "dynamic") || taskData.(map[string]interface{})["dynamicFile"] != "" || taskData.(map[string]interface{})["dynamicRange"] != "" {
		err := executeDynamicTask(taskName, commands, wfc, timeout, silent, stop, taskData, dockerImage, dockerhive, mounts, inputs)
		if err != nil {
			return err
		}
		return nil
	}
	err := executeSingleTask(taskName, commands, wfc, timeout, silent, stop, dockerImage, dockerhive, mounts, inputs)
	if err != nil {
		return err
	}
	return nil
}

package executors

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/utils"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
)

func executeCMD(taskName string, command string, stdoutDir string, stderrDir string, shell string, timeout time.Duration, displayStdout bool) error {
	// If no timeout is specified, set a default timeout of 1 day
	if timeout == 0 {
		timeout = time.Duration(24) * time.Hour
	}

	// Create a context with the specified timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a command and set its attributes
	stdoutFile := ""
	stderrFile := ""
	if stdoutDir != "" || stderrDir != "" {
		stdoutFile = stdoutDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stdout"
		stderrFile = stderrDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stderr"
	}

	cmd, err := createCommand(command, shell, ctx, stdoutFile, stderrFile, displayStdout)
	if err != nil {
		return fmt.Errorf("failed to create command for task '%s': %v", taskName, err)
	}

	// Start the command and capture any errors
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command for task '%s': %v", taskName, err)
	}

	// Create a channel for receiving timeout signals
	done := make(chan error, 1)

	// Run the command in a goroutine and wait for it to finish
	go func() {
		done <- cmd.Wait()
	}()

	// Check for timeout or completion of the command
	select {
	case <-ctx.Done():
		// Timeout occurred
		return fmt.Errorf("timeout occurred while executing command for task '%s'", taskName)
	case err := <-done:
		// Command has finished
		if err != nil {
			// Return the error
			return fmt.Errorf("error executing command for task '%s': %v", taskName, err)
		}
	}

	// Command has finished successfully
	return nil
}

func ExecHelper(config *config.Config) {
	// Tasks are false by default
	taskStatus, whitelist := getTaskStatusandWhitelist(config)
	// Create a wait group for parallel tasks
	var wg sync.WaitGroup

	// Function to execute a task
	executeTaskFunc := func(taskName string, taskData interface{}) {
		visuals.PrintState("START", taskName, "")
		err := executeTask(config, taskName, taskData)
		if err == nil {
			visuals.PrintState("SUCCESS", taskName, "")
			taskStatus[taskName] = true
		}
	}

	// Execute tasks in parallel
	for _, taskData := range config.Tasks {
		checkRequirements(&taskData, whitelist)

		// **Change:** Wait for required tasks to finish before executing the current task
		for _, requiredTask := range taskData.Required {
			for !taskStatus[requiredTask] {
				// Wait for 1 second before checking again
				time.Sleep(time.Second)
			}
		}

		if taskData.Parallel {
			// Execute parallel tasks without waiting
			wg.Add(1)
			go func(name string, data interface{}) {
				executeTaskFunc(name, data)
				wg.Done()
			}(taskData.Name, GetTaskDataWithName(taskData.Name, config))
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
					executeTaskFunc(taskData.Name, GetTaskDataWithName(taskData.Name, config))
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

func executeSingleTask(taskName string, commands []string, wfc *config.WorkflowConfig, timeout time.Duration, silent bool, stop bool) error {
	visuals.PrintState("Task-Info", taskName, "Task is Static")
	visuals.PrintState("Static-Task: "+taskName, taskName, "Executing Task")
	err := executeCMD(taskName, strings.Join(commands, " && "), wfc.StdeoutDir, wfc.StderrDir, wfc.Shell, timeout, silent)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") {
			visuals.PrintState("TIMEOUT", taskName, "")
		} else {
			visuals.PrintState("ERROR", taskName, "")
		}
		if stop {

			// gologger.Fatal().TimeStamp().Str("TaskName", taskName).Msgf("Stop On Error!")
			visuals.PrintState("FETAL", taskName, "")
			return err
		}
		return err
	}
	return nil
}

func executeTask(config *config.Config, taskName string, taskData interface{}) error {
	commands := utils.GetInterfaceVal(taskData, "cmds").([]string)
	timeout := utils.GetTimeout(taskData)
	silent := utils.GetInterfaceVal(taskData, "silent").(bool)
	silent = !silent
	stop := utils.GetInterfaceVal(taskData, "stop").(bool)
	wfc := &config.WorkflowConfig

	taskType := utils.GetInterfaceVal(taskData, "type").(string)
	if strings.ToLower(taskType) == "dynamic" {
		err := executeDynamicTask(taskName, commands, wfc, timeout, silent, stop, taskData)
		if err != nil {
			return err
		}
		return nil
	}
	err := executeSingleTask(taskName, commands, wfc, timeout, silent, stop)
	if err != nil {
		return err
	}
	return nil
}

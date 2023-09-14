package executors

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/utils"

	"github.com/projectdiscovery/gologger"
)

func executeCMD(taskName string, command string, stdoutDir string, stderrDir string, shell string, timeout time.Duration, displayStdout bool) error {
	// If no timeout is specified, set a default timeout of 1 day
	if timeout == 0 {
		timeout = time.Duration(24) * time.Hour
	}

	// Create a context with the specified timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// // Create a command and set its attributes
	// stdoutFile = stdoutFile + "/" + strings.ReplaceAll(stdoutFile, " ", "-") + ".stdout"
	// stderrFile = stdoutFile + "/" + strings.ReplaceAll(stdoutFile, " ", "-") + ".stderr"
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
			// If displayStdout is true, return the error
			if displayStdout {
				return fmt.Errorf("error executing command for task '%s': %v", taskName, err)
			}
			// Otherwise, return nil to indicate successful execution with no error
			return nil
		}
	}

	// Command has finished successfully
	return nil
}

func ExecHelper(config *config.Config) {
	requiredValues, err := GetRequiredValues(config)
	if err != nil {
		gologger.Fatal().TimeStamp().Msgf("Error %v", err)
		return
	}

	// Create a map to track the execution status of tasks
	taskStatus := make(map[string]bool)
	for taskName := range requiredValues {
		taskStatus[taskName] = false
	}

	// Create a wait group for parallel tasks
	var wg sync.WaitGroup

	// Function to execute a task
	executeTaskFunc := func(taskName string, taskData interface{}) {
		fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
		gologger.Warning().Label("Start").TimeStamp().Str("TaskName", taskName).Msg("Task Started")
		fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
		err := executeTask(config, taskName, taskData)
		if err == nil {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
			gologger.Warning().Label("Success").TimeStamp().Str("TaskName", taskName).Msg("Task Finished")
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
		}
		// Mark the task as executed
		taskStatus[taskName] = true
	}

	// Execute tasks in parallel
	for taskName, taskData := range requiredValues {
		if len((taskData.(map[string]interface{})["required"]).([]string)) != 0 {
			// parallel  false
			taskData.(map[string]interface{})["parallel"] = false
		}
		if utils.GetInterfaceVal(taskData, "parallel").(bool) {

			// Execute parallel tasks without waiting
			wg.Add(1)
			go func(name string, data interface{}) {
				executeTaskFunc(name, data)
				wg.Done()
			}(taskName, taskData)
		} else {
			// Function to execute tasks with dependencies
			executeWithDependencies := func() {
				taskData := requiredValues[taskName]
				requiredTasks := utils.GetInterfaceVal(taskData, "required").([]string)

				// Check if all required tasks are completed
				allRequirementsMet := true
				for _, requiredTask := range requiredTasks {
					if !taskStatus[requiredTask] {
						allRequirementsMet = false
						break
					}
				}

				// Execute the current task if all requirements are met
				if allRequirementsMet {
					executeTaskFunc(taskName, taskData)
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
	gologger.Info().Label("Task-Info").Str("TaskName", taskName).Msg("Task is Static")
	gologger.Debug().Label("Static-Task: " + taskName).TimeStamp().Msg("Executing Task")
	err := executeCMD(taskName, strings.Join(commands, " && "), wfc.StdeoutDir, wfc.StderrDir, wfc.Shell, timeout, silent)
	if err != nil {
		if stop {
			gologger.Fatal().TimeStamp().Str("TaskName", taskName).Msgf("Error executing task:")
			return err
		}
		gologger.Error().TimeStamp().Str("TaskName", taskName).Msgf("Error executing task:")
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
		return err
	}
	err := executeSingleTask(taskName, commands, wfc, timeout, silent, stop)
	return err
}

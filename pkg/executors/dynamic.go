package executors

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/utils"

	"github.com/projectdiscovery/gologger"
)

func executeDynamicTask(taskName string, commands []string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, taskData interface{}) error {
	dynamicFile, dynamicRange := "", ""
	maxThreads := 1 // Default to 1 thread if not specified in YAML

	if data, ok := taskData.(map[string]interface{}); ok {
		dynamicFile, _ = data["dynamicFile"].(string)
		dynamicRange, _ = data["dynamicRange"].(string)
		if maxThreadsYAML, ok := data["maxThreads"].(int); ok && maxThreadsYAML > 0 {
			maxThreads = maxThreadsYAML
		}
	}

	if (dynamicFile == "" && dynamicRange == "") || (dynamicFile != "" && dynamicRange != "") {
		gologger.Error().Label("ERROR").Str("TaskName", taskName).Msg("Either 'dynamicFile' or 'dynamicRange' must be specified, but not both.")
		if stop {
			gologger.Fatal().Label("STOP").Str("TaskName", taskName).Msg("Stop On Error!")
		}
		return nil
	}

	gologger.Info().Label("Task-Info: " + taskName).Msg("Task is Dynamic")
	if maxThreads > 1 {
		gologger.Info().Label("Task-Info: "+taskName).Str("Threads", strconv.Itoa(maxThreads)).Msgf("Running Tasks Parallel")
	}
	mergedcmd := strings.Join(commands, " && ")

	// Create a wait group to wait for all dynamic tasks to finish
	var wg sync.WaitGroup

	// Create a channel to signal task completion
	taskDone := make(chan struct{}, 2) // Two tasks: dynamic file and dynamic range

	// Execute dynamic file task concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		if dynamicFile != "" {
			lines, err := utils.ReadLinesFromFile(dynamicFile)
			if err != nil {
				gologger.Error().Label("ERROR").Str("TaskName", taskName).Msgf("Error Reading Lines from File: %v", err)
				if stop {
					gologger.Fatal().Label("STOP").Str("TaskName", taskName).Msg("Stop On Error!")
				}
				taskDone <- struct{}{}
				return
			}

			// Calculate the number of goroutines to spawn based on maxThreads
			numGoroutines := len(lines)
			if maxThreads < numGoroutines {
				numGoroutines = maxThreads
			}

			// Create a worker pool for executing dynamic file tasks concurrently
			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go dynamicWorker(taskName, mergedcmd, wfc, timeout, silent, stop, dynamicFile, lines[i*len(lines)/numGoroutines:(i+1)*len(lines)/numGoroutines], &wg, taskDone)
			}
		}
		taskDone <- struct{}{}
	}()

	// Execute dynamic range task concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		if dynamicRange != "" {
			// if dynamicRange contains more then 2 values
			ranges, err := utils.ConvertStringListToIntList(strings.Split(dynamicRange, ","))
			counter := utils.GenerateIntegerList(ranges[0], ranges[1])
			if len(counter) < maxThreads {
				maxThreads = len(counter)
				gologger.Info().Label("Task-Info: "+taskName).Str("Threads", strconv.Itoa(maxThreads)).Msgf("Threads > Range, Reducing...")

			}
			if err != nil {
				gologger.Error().Label("ERROR").Str("TaskName", taskName).Msgf("Invelid Range Format %s (expected: 1,5)", dynamicRange)
				if stop {
					gologger.Fatal().Label("STOP").Str("TaskName", taskName).Msg("Stop On Error!")
				}
				taskDone <- struct{}{}
				return
			}

			// Calculate the number of values per goroutine
			valuesPerGoroutine := (ranges[1] - ranges[0] + 1) / maxThreads

			// Create a worker pool for executing dynamic range tasks concurrently
			for i := 0; i < maxThreads; i++ {
				wg.Add(1)
				startIdx := ranges[0] + (i * valuesPerGoroutine)
				endIdx := startIdx + valuesPerGoroutine - 1
				go dynamicRangeWorker(taskName, mergedcmd, wfc, timeout, silent, stop, startIdx, endIdx, &wg, taskDone)
			}
		}
		taskDone <- struct{}{}
	}()

	// Wait for all dynamic tasks to finish
	go func() {
		wg.Wait()
		close(taskDone)
	}()

	// Block until all dynamic tasks are completed
	for range taskDone {
	}

	return nil
}

func dynamicWorker(taskName, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, dynamicFile string, lines []string, wg *sync.WaitGroup, taskDone chan<- struct{}) {
	defer wg.Done()
	for _, line := range lines {
		gologger.Debug().Label("Dynamic-Task: "+taskName).TimeStamp().Str("Value", line).Msg("Executing Task with FILE")
		dynamiccmd := strings.ReplaceAll(mergedcmd, "{{dynamicFile}}", line)
		newtaskName := taskName + "-" + line
		if err := executeCMD(newtaskName, dynamiccmd, wfc.StdeoutDir, wfc.StderrDir, wfc.Shell, timeout, silent); err != nil {
			if strings.Contains(err.Error(), "timeout") {
				gologger.Error().Label("TIMEOUT").Str("TaskName", taskName).Msgf("Error Executing Task: %v", err)
			} else {
				gologger.Error().Label("ERROR").Str("TaskName", taskName).Msgf("Error Executing Task: %v", err)
			}
			if stop {
				gologger.Fatal().Label("STOP").Str("TaskName", taskName).Msg("Stop On Error!")
			}
		}
	}
	taskDone <- struct{}{}
}

func dynamicRangeWorker(taskName, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, startIdx, endIdx int, wg *sync.WaitGroup, taskDone chan<- struct{}) {
	defer wg.Done()
	for i := startIdx; i <= endIdx; i++ {
		gologger.Debug().Label("Dynamic-Task: "+taskName).TimeStamp().Str("Value", strconv.Itoa(i)).Msg("Executing Task with range")
		dynamiccmd := strings.ReplaceAll(mergedcmd, "{{dynamicRange}}", strconv.Itoa(i))
		newtaskName := taskName + "-" + strconv.Itoa(i)
		if err := executeCMD(newtaskName, dynamiccmd, wfc.StdeoutDir, wfc.StderrDir, wfc.Shell, timeout, silent); err != nil {
			if strings.Contains(err.Error(), "timeout") {
				gologger.Error().Label("TIMEOUT").Str("TaskName", taskName).Msgf("Error Executing Task: %v", err)
			} else {
				gologger.Error().Label("ERROR").Str("TaskName", taskName).Msgf("Error Executing Task: %v", err)
			}
			if stop {
				gologger.Fatal().Label("STOP").Str("TaskName", taskName).Msg("Stop On Error!")
			}
		}
	}
	taskDone <- struct{}{}
}

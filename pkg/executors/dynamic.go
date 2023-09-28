package executors

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/utils"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
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
		visuals.PrintState("ERROR", taskName, "Either 'dynamicFile' or 'dynamicRange' must be specified, but not both.")
		if stop {
			visuals.PrintState("FETAL", taskName, "")
		}
		return nil
	}
	visuals.PrintState("Task-Info", taskName, "Task is Dynamic")
	if maxThreads > 1 {
		visuals.PrintStateDynamic("Task-Info", taskName, "Running Tasks Parallel", "Threads", strconv.Itoa(maxThreads))
	}
	mergedcmd := strings.Join(commands, " && ")

	var wg sync.WaitGroup

	taskDone := make(chan struct{}, 2) // Two tasks: dynamic file and dynamic range

	wg.Add(1)
	go func() {
		defer wg.Done()
		if dynamicFile != "" {
			lines, err := utils.ReadLinesFromFile(dynamicFile)
			if err != nil {
				visuals.PrintState("ERROR", taskName, "Error Reading Lines from File: "+err.Error())
				if stop {
					visuals.PrintState("FETAL", taskName, "")
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
				visuals.PrintStateDynamic("Task-Info", taskName, "Threads > Range, Reducing...", "Threads", strconv.Itoa(maxThreads))

			}
			if err != nil {
				visuals.PrintState("ERROR", taskName, "Invalid Range Format "+dynamicRange+" (expected: 1,5)")

				if stop {
					visuals.PrintState("FETAL", taskName, "")

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
		lineNumber := strconv.Itoa(utils.FindLineNumber(dynamicFile, line))
		visuals.PrintStateDynamic("Dynamic-Task: "+taskName, taskName, "Running Tasks Parallel", "FileLine", lineNumber)
		dynamiccmd := strings.ReplaceAll(mergedcmd, "{{dynamicFile}}", line)
		dynamiccmd = strings.ReplaceAll(dynamiccmd, "{{rand}}", utils.RandomString(5))
		dynamiccmd = strings.ReplaceAll(dynamiccmd, "{{dynamicOut}}", lineNumber)
		newtaskName := taskName + "-" + lineNumber
		if err := executeCMD(newtaskName, dynamiccmd, wfc.StdeoutDir, wfc.StderrDir, wfc.Shell, timeout, silent); err != nil {
			if strings.Contains(err.Error(), "timeout") {
				visuals.PrintState("TIMEOUT", taskName, "")
			} else {
				visuals.PrintState("ERROR", taskName, "")
			}
			if stop {
				visuals.PrintState("FETAL", taskName, "")
			}
		}
	}
	taskDone <- struct{}{}
}

func dynamicRangeWorker(taskName, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, startIdx, endIdx int, wg *sync.WaitGroup, taskDone chan<- struct{}) {
	defer wg.Done()
	for i := startIdx; i <= endIdx; i++ {
		visuals.PrintStateDynamic("Dynamic-Task: "+taskName, taskName, "Running Tasks Parallel", "Value", strconv.Itoa(i))
		dynamiccmd := strings.ReplaceAll(mergedcmd, "{{dynamicRange}}", strconv.Itoa(i))
		dynamiccmd = strings.ReplaceAll(dynamiccmd, "{{rand}}", utils.RandomString(10))
		dynamiccmd = strings.ReplaceAll(dynamiccmd, "{{dynamicOut}}", strconv.Itoa(i))
		newtaskName := taskName + "-" + strconv.Itoa(i)
		if err := executeCMD(newtaskName, dynamiccmd, wfc.StdeoutDir, wfc.StderrDir, wfc.Shell, timeout, silent); err != nil {
			if strings.Contains(err.Error(), "timeout") {
				visuals.PrintState("TIMEOUT", newtaskName, "")
			} else {
				visuals.PrintState("ERROR", newtaskName, "")
			}
			if stop {
				visuals.PrintState("FETAL", newtaskName, "")
			}
		}
	}
	taskDone <- struct{}{}
}

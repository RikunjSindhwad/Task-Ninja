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

func executeDynamicTask(taskName string, commands []string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, taskData interface{}, dockerimage string, dockerHive string, mounts, inputs []string) error {
	taskName = utils.SanitizeTaskName(taskName)
	defaultHive := wfc.DefaultHive
	if defaultHive == "" {
		defaultHive = "hive"
	}
	dynamicFile, dynamicRange := "", ""
	maxThreads := 1 // Default to 1 thread if not specified in YAML

	if data, ok := taskData.(map[string]interface{}); ok {
		dynamicFile, _ = data["dynamicFile"].(string)
		dynamicRange, _ = data["dynamicRange"].(string)
		dynamicFile = utils.ReplaceTaskPlaceholders(dynamicFile, defaultHive, "dynamic")
		dynamicRange = utils.ReplaceTaskPlaceholders(dynamicRange, defaultHive, "dynamic")
		// check if dynamicRange has filepath and file exists and if so read the file and replace the dynamicRange with the file content
		if strings.Contains(dynamicRange, "/") {
			dynamicRange = utils.ReadDynamicRangeFromFile(dynamicRange)
		}

		if maxThreadsYAML, ok := data["maxThreads"].(int); ok && maxThreadsYAML > 0 {
			maxThreads = maxThreadsYAML
		}
	}

	if (dynamicFile == "" && dynamicRange == "") || (dynamicFile != "" && dynamicRange != "") {
		visuals.PrintState("ERROR", taskName, "Either 'dynamicFile' or 'dynamicRange' must be specified, but not both.")
		if stop {
			visuals.PrintState("FATAL", taskName, "")
		}
		return nil
	}
	visuals.PrintState("Task-Info", taskName, "Task is Dynamic")
	if maxThreads > 1 {
		visuals.PrintStateDynamic("Task-Info", taskName, "Running Tasks Parallel", "Threads", strconv.Itoa(maxThreads))
	}
	mergedcmd := strings.Join(commands, " && ")

	var wg sync.WaitGroup

	taskDone := make(chan struct{}, maxThreads) // Adjusted channel size

	// Process dynamic file task concurrently
	if dynamicFile != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processDynamicFile(taskName, dynamicFile, mergedcmd, wfc, timeout, silent, stop, &wg, taskDone, dockerimage, dockerHive, maxThreads, mounts, inputs)
		}()
	}

	// Process dynamic range task concurrently
	if dynamicRange != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processDynamicRange(taskName, dynamicRange, mergedcmd, wfc, timeout, silent, stop, &wg, taskDone, dockerimage, dockerHive, maxThreads, mounts, inputs)
		}()
	}

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

// Separated the dynamic file processing into its own function for clarity.
func processDynamicFile(taskName, dynamicFile, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, wg *sync.WaitGroup, taskDone chan<- struct{}, dockerimage, dockerHive string, maxThreads int, mounts, inputs []string) {
	lines, err := utils.ReadLinesFromFile(dynamicFile)
	if err != nil {
		// ... error handling
		return
	}

	numGoroutines := len(lines)
	if maxThreads < numGoroutines {
		numGoroutines = maxThreads
	}

	for i := 0; i < numGoroutines; i++ {
		startIdx := i * len(lines) / numGoroutines
		endIdx := (i + 1) * len(lines) / numGoroutines
		if endIdx > len(lines) {
			endIdx = len(lines)
		}

		var linesWithNumbers []LineWithNumber
		for j, line := range lines[startIdx:endIdx] {
			linesWithNumbers = append(linesWithNumbers, LineWithNumber{Line: line, Number: startIdx + j + 1})
		}

		wg.Add(1)
		go dynamicWorker(taskName, mergedcmd, wfc, timeout, silent, stop, dynamicFile, linesWithNumbers, wg, taskDone, dockerimage, dockerHive, mounts, inputs)
	}
	taskDone <- struct{}{}
}

type LineWithNumber struct {
	Line   string
	Number int
}

// Separated the dynamic range processing into its own function for clarity.
func processDynamicRange(taskName, dynamicRange, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, wg *sync.WaitGroup, taskDone chan<- struct{}, dockerimage, dockerHive string, maxThreads int, mounts, inputs []string) {
	var ranges []int
	var err error
	// if dynamicRange contains more then 2 values
	if strings.Contains(dynamicRange, ",") {
		ranges, err = utils.ConvertStringListToIntList(strings.Split(dynamicRange, ","))
	}
	if strings.Contains(dynamicRange, "-") {
		ranges, err = utils.ConvertStringListToIntList(strings.Split(dynamicRange, "-"))
	}
	// check if range is int or not
	if err != nil {
		visuals.PrintState("ERROR", taskName, "Invalid Range Format "+dynamicRange+" (expected: 1,5 || 1-5)")
		if stop {
			visuals.PrintState("FATAL", taskName, "")
		}
		return
	}
	counter := utils.GenerateIntegerList(ranges[0], ranges[1])

	if len(counter) < maxThreads {
		maxThreads = len(counter)
		visuals.PrintStateDynamic("Task-Info", taskName, "Threads > Range, Reducing...", "Threads", strconv.Itoa(maxThreads))

	}
	if err != nil {
		visuals.PrintState("ERROR", taskName, "Invalid Range Format "+dynamicRange+" (expected: 1,5 || 1-5)")

		if stop {
			visuals.PrintState("FATAL", taskName, "") // Corrected the spelling here

		}
		taskDone <- struct{}{}
		return
	}

	// Calculate the number of values per goroutine
	valuesPerGoroutine := (ranges[1] - ranges[0] + 1) / maxThreads
	extraValues := (ranges[1] - ranges[0] + 1) % maxThreads

	currentStartIdx := ranges[0]
	// Create a worker pool for executing dynamic range tasks concurrently
	for i := 0; i < maxThreads; i++ {
		extra := 0
		if i < extraValues {
			extra = 1
		}

		startIdx := currentStartIdx
		endIdx := startIdx + valuesPerGoroutine + extra - 1
		currentStartIdx = endIdx + 1

		wg.Add(1)
		go dynamicRangeWorker(taskName, mergedcmd, wfc, timeout, silent, stop, startIdx, endIdx, wg, taskDone, dockerimage, dockerHive, mounts, inputs)
	}
	taskDone <- struct{}{}
}

// dynamicWorker and dynamicRangeWorker functions remain unchanged.

func dynamicWorker(taskName, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, dynamicFile string, lines []LineWithNumber, wg *sync.WaitGroup, taskDone chan<- struct{}, dockerimage string, dockerHive string, mounts, inputs []string) {
	defer wg.Done()
	for _, lineWithNumber := range lines {
		lineNumber := strconv.Itoa(lineWithNumber.Number)
		visuals.PrintStateDynamic("Dynamic-Task: "+taskName, taskName, "Running Tasks Parallel", "FileLine", lineNumber)
		dynamiccmd := strings.ReplaceAll(mergedcmd, "{{dynamicFile}}", lineWithNumber.Line)
		dynamiccmd = strings.ReplaceAll(dynamiccmd, "{{rand}}", utils.RandomString(5))
		newtaskName := taskName + ":" + lineNumber

		err := executeDockerCMD(newtaskName, dynamiccmd, wfc.DefaultHive, dockerHive, dockerimage, mounts, inputs, timeout, silent, true, wfc.EnableLogs)

		if err != nil {
			if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "waiting for container") {
				visuals.PrintState("TIMEOUT", taskName, err.Error())
			} else {
				visuals.PrintState("ERROR", taskName, err.Error())
			}
			if stop {
				visuals.PrintState("FATAL", taskName, err.Error())
			}
		}
	}
	taskDone <- struct{}{}
}

func dynamicRangeWorker(taskName, mergedcmd string, wfc *config.WorkflowConfig, timeout time.Duration, silent, stop bool, startIdx, endIdx int, wg *sync.WaitGroup, taskDone chan<- struct{}, dockerimage string, dockerHive string, mounts, inputs []string) {
	defer wg.Done()
	for i := startIdx; i <= endIdx; i++ {
		visuals.PrintStateDynamic("Dynamic-Task: "+taskName, taskName, "Running Tasks Parallel", "Value", strconv.Itoa(i))
		dynamiccmd := strings.ReplaceAll(mergedcmd, "{{dynamicRange}}", strconv.Itoa(i))
		dynamiccmd = strings.ReplaceAll(dynamiccmd, "{{rand}}", utils.RandomString(10))
		newtaskName := taskName + ":" + strconv.Itoa(i)

		err := executeDockerCMD(newtaskName, dynamiccmd, wfc.DefaultHive, dockerHive, dockerimage, mounts, inputs, timeout, silent, true, wfc.EnableLogs)

		if err != nil {
			if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "waiting for container") {
				visuals.PrintState("TIMEOUT", taskName, err.Error())
			} else {
				visuals.PrintState("ERROR", taskName, err.Error())
			}
			if stop {
				visuals.PrintState("FATAL", taskName, err.Error())
			}
		}
	}
	taskDone <- struct{}{}
}

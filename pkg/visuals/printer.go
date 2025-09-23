package visuals

import (
	"fmt"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
)

func PrintState(state, taskName, msg string) {

	switch strings.ToLower(state) {
	case "start":
		gologger.Info().Label(state).TimeStamp().Str("TaskName", taskName).Msg("Task Started")
	case "success":
		gologger.Info().Label(state).TimeStamp().Str("TaskName", taskName).Msg("Task Finished")
	case "timeout":
		gologger.Error().Label(state).Str("TaskName", taskName).Msg("Timeout occurred while executing task")
	case "error":
		gologger.Error().Label(state).Str("TaskName", taskName).TimeStamp().Msgf("%s", msg)
	case "task-info":
		gologger.Warning().Label(state).Str("TaskName", taskName).Msg(msg)
	case "static-task: " + strings.ToLower(taskName):
		gologger.Debug().Label(state).TimeStamp().Msg(msg)
	case "fatal":
		gologger.Fatal().TimeStamp().Str("TaskName", taskName).Msgf("Stop On Error!\n%s", msg)
	case "workflow-error":
		gologger.Fatal().Label("Workflow-Error").Msgf("%s", msg)

	}

}

func PrintStateDynamic(state, taskName, msg, str, value string) {

	switch strings.ToLower(state) {
	case "dynamic-task: " + strings.ToLower(taskName):
		gologger.Debug().TimeStamp().Label(state).Str(str, value).Msgf("%s", msg)
	case "task-info":
		gologger.Warning().Label(state).Str(str, value).Msg(msg)
	}
}

func PrintCredit(author, workflowName, wfType, duration string) {
	switch strings.ToLower(wfType) {
	case "start":
		fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
		gologger.Info().Label(PrintRandomColor("Workflow-Credit", 32)).Str("Workflow-Author", PrintRandomColor(author)).Msgf("Tasked Workflow '%s'", PrintRandomColor(workflowName))
		fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
	case "end":
		fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
		gologger.Info().Label(PrintRandomColor("Workflow-Complete", 32)).Str(PrintRandomColor("Time Taken"), PrintRandomColor(duration)).Msgf("Workflow '%s' Execution Complete", PrintRandomColor(workflowName))
		fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
	}

}

func SetLevelDebug() {

	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
}

func SetLevelWarning() {

	gologger.DefaultLogger.SetMaxLevel(levels.LevelWarning)
}

func SetLevelInfo() {

	gologger.DefaultLogger.SetMaxLevel(levels.LevelInfo)
}

func JsonView() {
	gologger.DefaultLogger.SetFormatter(&formatter.JSON{})
}

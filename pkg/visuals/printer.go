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
		fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
		gologger.Warning().Label(state).TimeStamp().Str("TaskName", taskName).Msg("Task Started")
		fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
	case "success":
		fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
		gologger.Warning().Label(state).TimeStamp().Str("TaskName", taskName).Msg("Task Finished")
		fmt.Fprintln(os.Stderr, strings.Repeat("-", 120))
	case "timeout":
		gologger.Error().Label(state).Str("TaskName", taskName).Msg("Timeout occurred while executing task")
	case "error":
		gologger.Error().Label(state).Str("TaskName", taskName).TimeStamp().Msgf("Error executing task")
	case "task-info":
		gologger.Info().Label(state).Str("TaskName", taskName).Msg(msg)
	case "static-task: " + strings.ToLower(taskName):
		gologger.Debug().Label(state).TimeStamp().Msg(msg)
	case "fetal":
		gologger.Fatal().TimeStamp().Str("TaskName", taskName).Msgf("Stop On Error!")
	case "workflow-error":
		gologger.Fatal().Label("Workflow-Error").Msgf(msg)

	}

}

func PrintStateDynamic(state, taskName, msg, str, value string) {

	switch strings.ToLower(state) {
	case "dynamic-task: " + strings.ToLower(taskName):
		gologger.Debug().TimeStamp().Label(state).Str(str, value).Msgf(msg)
	case "task-info":
		gologger.Info().Label(state).Str(str, value).Msg(msg)
	}
}

func PrintCredit(Author, workflowName string, Type string) {
	switch strings.ToLower(Type) {
	case "start":
		gologger.Info().Label("Workflow-Credit").Str("Workflow-Author", Author).Msgf("Tasked Workflow '%s'", workflowName)
		fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
	case "end":
		gologger.Info().Label("Workflow-Complete").Str("Workflow-Author", Author).Msgf("Workflow '%s' Execution Complete", workflowName)
		fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
	}

}

func SetLevelDebug() {

	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
}

func SetLevelWarning() {

	gologger.DefaultLogger.SetMaxLevel(levels.LevelWarning)
}

func JsonView() {
	gologger.DefaultLogger.SetFormatter(&formatter.JSON{})
}

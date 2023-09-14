package executors

import (
	"context"
	"os"
	"os/exec"
	"runtime"

	"github.com/projectdiscovery/gologger"
)

func createCommand(command, shell string, ctx context.Context, stdoutFile, stderrFile string, displayStdout bool) (*exec.Cmd, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		if shell == "" {
			shell = "cmd"
		}
		cmd = exec.CommandContext(ctx, shell, "/C", command)
	default:
		if shell == "" {
			shell = "sh"
		}
		cmd = exec.CommandContext(ctx, shell, "-c", command)
	}

	// Create or open files for logging stdout and stderr if specified
	var stdoutWriter, stderrWriter *os.File
	var err error

	if stdoutFile != "" {
		stdoutWriter, err = os.OpenFile(stdoutFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err == nil {
			cmd.Stdout = stdoutWriter
		} else {
			gologger.Error().Label("ERROR").Msgf("Error Opening stdout File: %v", err)
		}

	}

	if stderrFile != "" {
		stderrWriter, err = os.OpenFile(stderrFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err == nil {
			cmd.Stderr = stderrWriter
		} else {
			gologger.Error().Label("ERROR").Msgf("Error Opening stderr File: %v", err)
		}

	}

	if displayStdout {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}

	return cmd, nil
}

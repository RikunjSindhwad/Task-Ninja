package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
)

func ReplacePlaceholders(config *config.Config) {

	// Replace placeholders in the Tasks slice
	for i := range config.Tasks {

		// Replace placeholders in the dynamicFile within each Task
		if config.Tasks[i].DynamicFile != "" {
			config.Tasks[i].DynamicFile = ReplacePlaceholdersInString(config.Tasks[i].DynamicFile, config.Vars)
		}

		for input := range config.Tasks[i].Inputs {
			config.Tasks[i].Inputs[input] = ReplacePlaceholdersInString(config.Tasks[i].Inputs[input], config.Vars)
		}

		// Replace placeholders in the Cmds slice within each Task
		for j := range config.Tasks[i].Cmds {
			config.Tasks[i].Cmds[j] = ReplacePlaceholdersInString(config.Tasks[i].Cmds[j], config.Vars)
		}
	}
}

// Function to replace placeholders in a string
func ReplacePlaceholdersInString(input string, vars map[string]string) string {
	for key, value := range vars {
		placeholder := fmt.Sprintf("{{%s}}", key)
		input = strings.ReplaceAll(input, placeholder, value)
	}
	return input
}

func ReplaceDockerplaceholders(command, dockerHive, dockerHiveTaskDir, hostHiveTaskDir, hostHiveTaskOutputDir, hostHiveTaskInputDir string) string {
	command = strings.ReplaceAll(command, "{{hiveout}}", filepath.Join(dockerHive, "out"))
	command = strings.ReplaceAll(command, "{{hivein}}", filepath.Join(dockerHive, "in"))
	command = strings.ReplaceAll(command, "{{hive}}", dockerHive)
	command = strings.ReplaceAll(command, "{{hosthive}}", hostHiveTaskDir)
	command = strings.ReplaceAll(command, "{{hosthiveout}}", hostHiveTaskOutputDir)
	command = strings.ReplaceAll(command, "{{hosthivein}}", hostHiveTaskInputDir)
	return command
}

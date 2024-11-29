package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/v2/pkg/config"
)

func ReplacePlaceholders(configuration *config.Config) {

	// Replace placeholders in the Tasks slice
	for i := range configuration.Tasks {

		// Replace placeholders in the dynamicFile within each Task
		if configuration.Tasks[i].DynamicFile != "" {
			configuration.Tasks[i].DynamicFile = ReplacePlaceholdersInString(configuration.Tasks[i].DynamicFile, configuration.Vars)
		}

		for input := range configuration.Tasks[i].Inputs {
			configuration.Tasks[i].Inputs[input] = ReplacePlaceholdersInString(configuration.Tasks[i].Inputs[input], configuration.Vars)
		}

		// Replace placeholders in the Cmds slice within each Task
		for j := range configuration.Tasks[i].Cmds {
			configuration.Tasks[i].Cmds[j] = ReplacePlaceholdersInString(configuration.Tasks[i].Cmds[j], configuration.Vars)
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

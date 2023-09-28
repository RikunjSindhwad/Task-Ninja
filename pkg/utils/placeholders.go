package utils

import (
	"fmt"
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
)

func ReplacePlaceholders(config *config.Config) {

	// Replace placeholders in the Tasks slice
	for i := range config.Tasks {

		// Replace placeholders in the dynamicFile within each Task
		config.Tasks[i].DynamicFile = ReplacePlaceholdersInString(config.Tasks[i].DynamicFile, config.Vars)

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

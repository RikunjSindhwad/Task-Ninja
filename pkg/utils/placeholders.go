package utils

import (
	"Robensive-TaskNinja/pkg/config"
	"fmt"
	"strings"
)

func ReplacePlaceholders(config *config.Config) {
	// Replace placeholders in the Usage and Author fields
	config.WorkflowConfig.Usage = ReplacePlaceholdersInString(config.WorkflowConfig.Usage, config.Vars)
	config.WorkflowConfig.Author = ReplacePlaceholdersInString(config.WorkflowConfig.Author, config.Vars)

	// Replace placeholders in the Tasks slice
	for i := range config.Tasks {
		config.Tasks[i].Name = ReplacePlaceholdersInString(config.Tasks[i].Name, config.Vars)

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

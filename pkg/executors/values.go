package executors

import (
	"fmt"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
)

func GetRequiredValues(config *config.Config) (map[string]interface{}, error) {
	requiredValues := make(map[string]interface{})

	for _, task := range config.Tasks {
		if len(task.Required) > 0 {
			for _, requiredTaskName := range task.Required {
				found := false
				for _, checkTask := range config.Tasks {
					if checkTask.Name == requiredTaskName {
						found = true
						break
					}
				}
				if !found {
					return nil, fmt.Errorf("required task '%s' not found for task '%s'", requiredTaskName, task.Name)
				}
			}
		}

		requiredValues[task.Name] = map[string]interface{}{
			"cmds":         task.Cmds,
			"parallel":     task.Parallel,
			"timeout":      task.Timeout,
			"required":     task.Required,
			"silent":       task.Silent,
			"stop":         task.StoponError,
			"type":         task.Type,
			"dynamicFile":  task.DynamicFile,
			"dynamicRange": task.DynamicRange,
			"maxThreads":   task.MaxThreads,
		}
	}

	return requiredValues, nil
}

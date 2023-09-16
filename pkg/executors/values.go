package executors

import (
	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
)

func GetTaskDataWithName(taskName string, config *config.Config) map[string]interface{} {
	for _, task := range config.Tasks {
		if task.Name == taskName {
			taskData := map[string]interface{}{
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
			return taskData
		}
	}

	return nil
}

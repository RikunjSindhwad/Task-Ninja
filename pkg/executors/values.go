package executors

import (
	"github.com/RikunjSindhwad/Task-Ninja/v2/pkg/config"
)

func GetTaskDataWithName(taskName string, configuration *config.Config) map[string]interface{} {
	for i := range configuration.Tasks {
		task := &configuration.Tasks[i]
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
				"image":        task.Image,
				"dockerHive":   task.DockerHive,
				"inputMounts":  task.InputMouts,
				"inputs":       task.Inputs,
			}
			return taskData
		}
	}

	return nil
}

package executors

import (
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
)

func checkRequiredTasks(task *config.Task) bool {
	if len(task.Required) != 0 && task.Parallel {
		return true
	} else if len(task.Required) != 0 && !task.Parallel {
		return true

	}
	return false
}

func checkRequirements(taskData *config.Task, whitelist []string) {
	if checkRequiredTasks(taskData) {
		visuals.PrintState("Task-Info", taskData.Name, "There are required Tasks "+strings.Join(taskData.Required, ", "))
		taskData.Parallel = false
	}
	if checkListContainsSame(taskData.Required, taskData.Name) {
		visuals.PrintState("Task-Info", taskData.Name, "Removed Same task from required list")
		taskData.Required = removeItemFromList(taskData.Required, taskData.Name)
	}
	taskData.Required = checkListContainsUnknwn(taskData.Required, whitelist)
}
func getTaskStatusandWhitelist(config *config.Config) (map[string]bool, []string) {
	taskStatus := make(map[string]bool)
	whitelist := []string{}
	for _, task := range config.Tasks {
		taskName := task.Name
		taskStatus[taskName] = false
		whitelist = append(whitelist, taskName)
	}
	return taskStatus, whitelist
}

func removeItemFromList(list []string, value string) []string {
	for i, item := range list {
		if item == value {
			list = append(list[:i], list[i+1:]...)
			return list
		}
	}
	return list
}
func checkListContainsSame(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func checkListContainsUnknwn(requiredList []string, whitelist []string) []string {
	updatedRequiredList := []string{}
	for _, item := range requiredList {
		if checkListContainsSame(whitelist, item) {
			updatedRequiredList = append(updatedRequiredList, item)
		}
	}
	return updatedRequiredList
}

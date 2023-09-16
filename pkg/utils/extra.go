package utils

import (
	"errors"
	"strconv"
	"time"
)

func ConvertStringListToIntList(strList []string) ([]int, error) {
	if len(strList) > 2 {
		return nil, errors.New("invalid range")
	}
	intList := make([]int, len(strList))
	for i, str := range strList {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err // Return an error if conversion fails
		}
		intList[i] = num
	}
	return intList, nil
}

func GenerateIntegerList(start, end int) []int {
	if start > end {
		return nil
	}

	list := GenerateIntegerList(start, end-1)
	list = append(list, end)
	return list
}

func GetTimeout(taskData interface{}) time.Duration {
	seconds := taskData.(map[string]interface{})["timeout"].(int)
	return time.Duration(seconds) * time.Second
}

func AllDependenciesSatisfied(taskData interface{}, taskStatus map[string]bool) bool {
	requiredTasks := GetInterfaceVal(taskData, "required").([]string)
	for _, requiredTask := range requiredTasks {
		if !taskStatus[requiredTask] {
			return false
		}
	}
	return true
}

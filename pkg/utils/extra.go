package utils

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsURL(str string) bool {
	re := regexp.MustCompile(`^https?://`)
	return re.MatchString(str)
}

func SanitizeTaskName(taskName string) string {
	// Define a regular expression to match special characters
	specialChars := regexp.MustCompile("[^a-zA-Z0-9_-]+")

	// Replace special characters with underscores
	sanitized := specialChars.ReplaceAllString(taskName, "-")

	return sanitized
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

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

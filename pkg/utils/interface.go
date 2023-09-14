package utils

func GetInterfaceVal(taskData interface{}, value string) interface{} {
	dataMap, ok := taskData.(map[string]interface{})
	if !ok {
		// Handle the case where taskData is not a map
		return nil
	}

	result := dataMap[value]
	return result
}

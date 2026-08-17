package task

import "fmt"

var tasksMap = make(map[string]Task)

// Helper Functions
func keyExists(key string) bool {
	_, exists := tasksMap[key]
	return exists
}

//task funcs

func addTask(taskKey string) bool {
	if keyExists(taskKey) {
		newTask := Task{Status: "pending"} //make new Task struct Instance
		tasksMap[taskKey] = newTask
		return true
	}
	return false
}

func deleteTask(taskKey string) bool {
	if keyExists(taskKey) {
		delete(tasksMap, taskKey)
		return true
	}
	return false
}

func printTask() {
	for index, value := range tasksMap {
		fmt.Printf("Task #%v : %v", index, value)
	}
}

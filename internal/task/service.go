package task

import (
	"fmt"
	"strings"
)

var tasksMap = make(map[string]*Task)

// Helper Functions
func keyExists(key string) bool {
	_, exists := tasksMap[key]
	return exists
}

//task funcs

func addTask(taskKey string) bool {
	if !keyExists(taskKey) {
		newTask := &Task{Status: "pending"} //make new Task struct Instance
		tasksMap[taskKey] = newTask
		saveDataToFile(tasksMap)
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
	count := 0
	for index, value := range tasksMap {
		fmt.Printf("Task #%d: [%v:%v], ", count, index, *value)
		count++
	}
}

func updateTask(taskKey string) bool {
	if keyExists(taskKey) {
		tasksMap[taskKey].Status = "Done"
		fmt.Printf("Task(%v) Updated\n", taskKey)
		return true
	}
	return false
}

func printTaskStatus(taskKey string) bool {
	if keyExists(taskKey) {
		fmt.Printf("Task Status: %v\n", tasksMap[taskKey].Status)
		return true
	}
	return false
}

func sliceToString(s []string) string {
	return strings.Join(s, " ")
}

func Run() {
	cliArgs := getArgs()
	//fmt.Println(cliArgs)
	switch cliArgs[0] {
	case "add":
		addTask(sliceToString(cliArgs[1:]))
	case "delete":
		deleteTask(sliceToString(cliArgs[1:]))
	case "list":
		printTask()
	default:
		fmt.Println("No Command Found!")
	}

}

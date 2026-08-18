package task

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveDataToFile(mp map[string]*Task) bool {
	jsonData, err := json.Marshal(mp)
	if err != nil {
		fmt.Println("Error Marshling: ", err)
		return false
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error Writing to json: ", err)
		return false
	}
	fmt.Println("Saved to tasks.json")
	return true
}

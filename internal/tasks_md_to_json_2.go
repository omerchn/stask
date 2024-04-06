package internal

import (
	"encoding/json"
	"os"
)

type Task2 struct {
	Task        string `json:"task"`
	Description string `json:"description"`
}

type Subject []Task2

type Category map[string]Subject

type Result2 map[string]Category

func TasksToJson2(tasksFile string) (string, error) {
	f, err := os.Open(tasksFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := Result2{}

	resultStr, err := json.MarshalIndent(r, "- ", "  ")
	if err != nil {
		return "", err
	}

	return string(resultStr), nil
}

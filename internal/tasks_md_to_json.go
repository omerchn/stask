package internal

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type Task struct {
	Task        string `json:"task"`
	Description string
	Category    string
	Subject     string
}

type Result struct {
	Tasks      []Task   `json:"tasks"`
	Categories []string `json:"categories"`
	Subjects   []string `json:"subjects"`
}

func TasksToJson(tasksFile string) (string, error) {
	f, err := os.Open(tasksFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := Result{}

	currentCategory := ""
	currentSubject := ""
	currentTask := ""
	var currentDescriptionLines []string

	addCurrentTask := func() {
		description := strings.Join(currentDescriptionLines, "\n")
		t := Task{Task: currentTask, Description: description, Category: currentCategory, Subject: currentSubject}
		r.Tasks = append(r.Tasks, t)
		currentTask = ""
		currentDescriptionLines = nil
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if currentTask != "" && strings.HasPrefix(line, descriptionPrefix) && len(line) > len(descriptionPrefix) {
			currentDescriptionLines = append(currentDescriptionLines, strings.TrimPrefix(line, descriptionPrefix))
			continue
		}
		if currentTask != "" {
			addCurrentTask()
		}
		if strings.HasPrefix(line, categoryPrefix) {
			category := strings.TrimPrefix(line, categoryPrefix)
			r.Categories = append(r.Categories, category)
			currentCategory = category
		}
		if strings.HasPrefix(line, subjectPrefix) {
			subject := strings.TrimPrefix(line, subjectPrefix)
			r.Subjects = append(r.Subjects, subject)
			currentSubject = subject
		}
		if strings.HasPrefix(line, taskPrefix) {
			currentTask = strings.TrimPrefix(line, taskPrefix)
		}
	}
	addCurrentTask()

	err = scanner.Err()
	if err != nil {
		return "", err
	}

	resultStr, err := json.MarshalIndent(r, "- ", "  ")
	if err != nil {
		return "", err
	}

	return string(resultStr), nil
}

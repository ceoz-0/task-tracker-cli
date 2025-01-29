package tasks

import "time"

func create(description string) Task {
	newTask := Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasksMapStorage[len(tasksMapStorage)+1] = newTask
	return newTask

}

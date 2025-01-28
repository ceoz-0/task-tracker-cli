package tasks

import "time"

func create(description string) task {
	newTask := task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasksMapStorage[len(tasksMapStorage)+1] = newTask
	return newTask

}

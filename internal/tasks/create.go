package tasks

import "time"

func create(description string) Task {
	newTask := Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	nextTaskId++
	tasksMapStorage[nextTaskId] = newTask
	return newTask

}

package tasks

import "time"

func updateDescription(id int, description string) {
	task := tasksMapStorage[id]
	task.Description = description
	task.UpdatedAt = time.Now()
	tasksMapStorage[id] = task
}

func updateStatus(id int, status string) {
	task := tasksMapStorage[id]
	task.Status = status
	task.UpdatedAt = time.Now()
	tasksMapStorage[id] = task
}

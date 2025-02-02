package tasks

import "time"

func update(id int, description string, status string) {

	task := tasksMapStorage[id]

	task.Description = description
	task.Status = status
	task.UpdatedAt = time.Now()
	tasksMapStorage[id] = task
}

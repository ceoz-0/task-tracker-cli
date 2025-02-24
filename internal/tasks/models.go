package tasks

import "time"

type Task struct {
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var tasksMapStorage = map[int]Task{}
var nextTaskId = 0

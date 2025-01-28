package tasks

import "time"

type task struct {
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var tasksMapStorage = map[int]task{}

package tasks

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	task := create("Test Task")

	if task.Description != "Test Task" {
		t.Errorf("Expected 'Test Task', got '%s' ", task.Description)
	}

	if task.Status != "todo" {
		t.Errorf("Expected 'todo', got '%s' ", task.Status)
	}

}

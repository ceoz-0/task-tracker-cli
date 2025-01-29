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

type TaskFilter struct {
	ID       *int   // ✅ Pointer allows filtering by ID or leaving it empty
	Status   string // ✅ Optional status filter
	Priority string // ✅ Optional priority filter
}

func TestReadTask(t *testing.T) {
	// ✅ Clear & Set Up Fake Data (Isolation)
	tasksMapStorage = make(map[int]Task)

	testTask := Task{Description: "Test Read", Status: "completed"}
	tasksMapStorage[1] = testTask // ✅ ID is the MAP KEY, NOT a field!

	// ✅ Call ReadTasks
	filter := TaskFilter{ID: new(int)} // ✅ Create a pointer to pass ID filter
	*filter.ID = 1                     // ✅ Set it to the map key

	result, found := read(filter)

	// ✅ Assertions
	if !found {
		t.Errorf("Expected to find task with ID 1, but got nothing")
	}

	if len(result) != 1 || result[0].Description != "Test Read" {
		t.Errorf("Expected 'Test Read', got %+v", result)
	}

	// ✅ Cleanup after test (reset storage)
	defer func() { tasksMapStorage = make(map[int]Task) }()
}

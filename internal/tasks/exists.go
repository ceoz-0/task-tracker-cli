package tasks

func exists(id int) bool {
	_, exists := tasksMapStorage[id]
	return exists

}

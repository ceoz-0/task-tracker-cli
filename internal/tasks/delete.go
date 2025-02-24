package tasks

func deleteTask(id int) {

	delete(tasksMapStorage, id)
	nextTaskId--

}

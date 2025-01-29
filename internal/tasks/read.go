package tasks

func read(tf TaskFilter) ([]Task, bool) {
	var results []Task

	if tf.ID != nil {
		if Task, exists := tasksMapStorage[*tf.ID]; exists {

			results = append(results, Task)
			return results, true

		}
		return results, false
	}
	for _, task := range tasksMapStorage {
		if tf.Status != "" && task.Status != tf.Status {
			continue
		}
		results = append(results, task)

	}
	return results, len(results) > 0

}

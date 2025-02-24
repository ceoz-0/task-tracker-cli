package tasks

func validStatus(stat string) bool {

	validStat := map[string]interface{}{
		"todo":        nil,
		"in-progress": nil,
		"done":        nil,
	}

	_, exists := validStat[stat]

	return exists

}

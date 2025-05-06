package utils

var Status = map[int]string{
	1: "Pending",
	2: "In Progress",
	3: "Completed",
	4: "Canceled",
}

func GetStatus(status int) string {
	if status >= 1 && status <= 4 {
		return Status[status]
	}

	return ""
}

var MaxHirarcy = 4

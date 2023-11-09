package taskdtoupdate

type TaskUpdateInputDTO struct {
	Name    string `json:"name"`
	DueTime string `json:"due_time"`
}

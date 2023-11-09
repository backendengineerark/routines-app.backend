package taskdtocreate

type TaskCreateInputDTO struct {
	Name    string `json:"name"`
	DueTime string `json:"due_time"`
}

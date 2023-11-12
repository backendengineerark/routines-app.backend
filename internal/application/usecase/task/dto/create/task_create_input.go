package taskdtocreate

import weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"

type TaskCreateInputDTO struct {
	Name    string                            `json:"name"`
	DueTime string                            `json:"due_time"`
	Days    []weekdaydtolist.WeekdayOutputDTO `json:"weekdays"`
}

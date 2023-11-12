package taskdtoupdate

import (
	"time"

	weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"
)

type TaskUpdateOutputDTO struct {
	Id         string                            `json:"id"`
	Name       string                            `json:"name"`
	DueTime    string                            `json:"due_time"`
	IsArchived bool                              `json:"is_archived"`
	CreatedAt  time.Time                         `json:"created_at"`
	UpdatedAt  time.Time                         `json:"updated_at"`
	Days       []weekdaydtolist.WeekdayOutputDTO `json:"weekdays"`
}

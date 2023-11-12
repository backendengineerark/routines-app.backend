package taskdtolist

import (
	"time"

	weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"
)

type TaskListOutputDTO struct {
	Id             string                            `json:"id"`
	Name           string                            `json:"name"`
	DueTime        string                            `json:"due_time"`
	IsArchived     bool                              `json:"is_archived"`
	CompletedTimes int64                             `json:"completed_times"`
	FailedTimes    int64                             `json:"failed_times"`
	Days           []weekdaydtolist.WeekdayOutputDTO `json:"weekdays"`
	CreatedAt      time.Time                         `json:"created_at"`
	UpdatedAt      time.Time                         `json:"updated_at"`
}

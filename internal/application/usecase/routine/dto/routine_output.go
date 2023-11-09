package routinedto

import (
	"time"

	taskdtolist "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/list"
)

type RoutineOutputDTO struct {
	Id            string                         `json:"id"`
	Task          *taskdtolist.TaskListOutputDTO `json:"task"`
	ReferenceDate time.Time                      `json:"reference_date"`
	IsFinished    bool                           `json:"is_finished"`
	CreatedAt     time.Time                      `json:"created_at"`
	UpdatedAt     time.Time                      `json:"updated_at"`
}

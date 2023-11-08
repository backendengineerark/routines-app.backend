package routinedto

import (
	"time"

	taskdto "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto"
)

type RoutineOutputDTO struct {
	Id            string                 `json:"id"`
	Task          *taskdto.TaskOutputDTO `json:"task"`
	ReferenceDate time.Time              `json:"reference_date"`
	IsFinished    bool                   `json:"is_finished"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

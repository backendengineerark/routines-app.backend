package taskdto

import "time"

type RoutineOutputDTO struct {
	Id            string         `json:"id"`
	Task          *TaskOutputDTO `json:"task"`
	ReferenceDate time.Time      `json:"reference_date"`
	IsFinished    bool           `json:"is_finished"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

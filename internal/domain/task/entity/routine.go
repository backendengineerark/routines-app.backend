package entity

import (
	"time"

	customerrors "github.com/backendengineerark/routines-app/internal/domain/common/custom_errors"
	"github.com/google/uuid"
)

type Routine struct {
	Id         string
	Task       *Task
	DueDate    time.Time
	IsFinished bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func CreateRoutine(task *Task, dueDate time.Time) (*Routine, error) {
	routine := &Routine{
		Id:         uuid.NewString(),
		Task:       task,
		DueDate:    dueDate,
		IsFinished: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := routine.IsValid()
	if err != nil {
		return nil, err
	}

	return routine, nil
}

func (r *Routine) IsValid() error {
	if r.DueDate.IsZero() {
		return &customerrors.BusinessValidationError{
			Message: "Due date is required to create a routine",
		}
	}
	return nil
}

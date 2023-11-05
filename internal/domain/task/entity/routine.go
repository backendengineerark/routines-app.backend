package entity

import (
	"time"

	customerrors "github.com/backendengineerark/routines-app/internal/domain/common/custom_errors"
	"github.com/google/uuid"
)

type Routine struct {
	Id            string
	Task          *Task
	ReferenceDate time.Time
	IsFinished    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func CreateRoutine(task *Task, referenceDate time.Time) (*Routine, error) {
	routine := &Routine{
		Id:            uuid.NewString(),
		Task:          task,
		ReferenceDate: referenceDate,
		IsFinished:    false,
		CreatedAt:     time.Now().Local(),
		UpdatedAt:     time.Now().Local(),
	}

	err := routine.IsValid()
	if err != nil {
		return nil, err
	}

	return routine, nil
}

func (r *Routine) IsValid() error {
	if r.ReferenceDate.IsZero() {
		return &customerrors.BusinessValidationError{
			Message: "Due date is required to create a routine",
		}
	}
	return nil
}

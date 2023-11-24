package entity

import (
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/common/custom_dates"
	customerrors "github.com/backendengineerark/routines-app/internal/domain/common/custom_errors"
	"github.com/google/uuid"
)

type Routine struct {
	Id            string
	ReferenceDate time.Time
	IsFinished    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func CreateRoutine() (*Routine, error) {
	routine := &Routine{
		Id:            uuid.NewString(),
		ReferenceDate: custom_dates.TodayBeginningHour(),
		IsFinished:    false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

func (r *Routine) Finish() {
	r.IsFinished = true
	r.UpdatedAt = time.Now()
}

func (r *Routine) Unfinish() {
	r.IsFinished = false
	r.UpdatedAt = time.Now()
}

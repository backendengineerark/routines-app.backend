package entity

import (
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/common/custom_dates"
	customerrors "github.com/backendengineerark/routines-app/internal/domain/common/custom_errors"
	"github.com/google/uuid"
)

type CreateTaskCommand struct {
	Name    string
	DueTime string
}

type Task struct {
	Id        string
	Name      string
	DueTime   string
	IsArchive bool
	Routines  []*Routine
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateTask(command *CreateTaskCommand) (*Task, error) {
	task := &Task{
		Id:        uuid.NewString(),
		Name:      command.Name,
		DueTime:   command.DueTime,
		IsArchive: false,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}
	task.AddRoutine(custom_dates.TodayBeginningHour())
	err := task.IsValid()

	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ta *Task) IsValid() error {
	if ta.Name == "" {
		return &customerrors.BusinessValidationError{
			Message: "Task name is required",
		}
	}

	if ta.DueTime == "" {
		return &customerrors.BusinessValidationError{
			Message: "Task due time is required",
		}
	}

	return nil
}

func (ta *Task) AddRoutine(dueTime time.Time) error {
	routine, err := CreateRoutine(ta, dueTime)
	if err != nil {
		return err
	}
	ta.Routines = append(ta.Routines, routine)

	return nil
}

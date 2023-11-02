package entity

import (
	"time"

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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateTask(command *CreateTaskCommand) (*Task, error) {
	task := &Task{
		Id:        uuid.NewString(),
		Name:      command.Name,
		DueTime:   command.DueTime,
		IsArchive: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
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

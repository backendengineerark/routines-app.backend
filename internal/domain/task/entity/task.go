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
	Id           string
	Name         string
	DueTime      string
	IsArchive    bool
	TodayRoutine *Routine
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	err := task.IsValid()
	if err != nil {
		return nil, err
	}

	routine, err := CreateRoutine()
	if err != nil {
		return nil, err
	}

	task.TodayRoutine = routine

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

func (ta *Task) Archive() {
	ta.IsArchive = true
}

func (ta *Task) Unarchive() {
	ta.IsArchive = false
}

func (ta *Task) CreateTodayRoutine() error {
	routine, err := CreateRoutine()
	if err != nil {
		return err
	}

	ta.TodayRoutine = routine

	return nil
}

func (ta *Task) FinishTodayRoutine() error {
	if ta.TodayRoutine == nil {
		return &customerrors.BusinessValidationError{
			Message: "No today routine registered to finish",
		}
	}
	ta.TodayRoutine.Finish()

	return nil
}

func (ta *Task) UnfinishTodayRoutine() error {
	if ta.TodayRoutine == nil {
		return &customerrors.BusinessValidationError{
			Message: "No today routine registered to finish",
		}
	}
	ta.TodayRoutine.Unfinish()

	return nil
}

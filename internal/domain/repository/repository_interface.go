package repository

import (
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type ITaskRepository interface {
	Create(task *entity.Task) error
	Update(task *entity.Task) error
	ListBy(isArchived bool) ([]entity.Task, error)
	ListRoutine(date time.Time) ([]entity.Routine, error)
}

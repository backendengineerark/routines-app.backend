package repository

import (
	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type ITaskRepository interface {
	Create(task *entity.Task) error
	FindById(id string) (*entity.Task, error)
	Update(task *entity.Task) error
	CreateTodayRoutine(task *entity.Task) error
	UpdateTodayRoutine(task *entity.Task) error
	FindAllBy(isArchived bool) ([]*entity.Task, error)
}

package repository

import (
	"time"

	metricentity "github.com/backendengineerark/routines-app/internal/domain/metric/entity"
	taskentity "github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type ITaskRepository interface {
	Create(task *taskentity.Task) error
	FindById(id string) (*taskentity.Task, error)
	Update(task *taskentity.Task) error
	CreateTodayRoutine(task *taskentity.Task) error
	UpdateTodayRoutine(task *taskentity.Task) error
	FindAllBy(isArchived bool) ([]*taskentity.Task, error)
}

type IMetricRepository interface {
	FindAllRoutinesByRangeDate(initial time.Time, end time.Time) ([]*metricentity.Metric, error)
}

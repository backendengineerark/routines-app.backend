package repository

import (
	"time"

	metricentity "github.com/backendengineerark/routines-app/internal/domain/metric/entity"
	taskentity "github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type ITaskRepository interface {
	Create(task *taskentity.Task) error
	FindById(id string) (*taskentity.Task, error)
	FindRoutineById(id string) (*taskentity.Routine, error)
	Update(task *taskentity.Task) error
	Delete(task *taskentity.Task) error
	CreateTodayRoutine(task *taskentity.Task) error
	UpdateTodayRoutine(task *taskentity.Task) error
	FindAllBy(isArchived bool) ([]*taskentity.Task, error)
	FindAllWeekday() ([]*taskentity.Weekday, error)
	FindWeekdayIn(weekdayIds []string) ([]*taskentity.Weekday, error)
	DeleteRoutine(routine *taskentity.Routine) error
}

type IMetricRepository interface {
	FindAllRoutinesByRangeDate(initial time.Time, end time.Time) ([]*metricentity.Metric, error)
}

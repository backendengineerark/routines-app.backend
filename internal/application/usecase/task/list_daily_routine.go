package usecase

import (
	"sort"
	"time"

	taskdto "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type ListRoutineUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewListRoutineUseCase(taskRepository repository.ITaskRepository) *ListRoutineUseCase {
	return &ListRoutineUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *ListRoutineUseCase) Execute(date time.Time) ([]taskdto.RoutineOutputDTO, error) {

	isArchived := false
	tasks, err := ct.TaskRepository.FindAllBy(isArchived)

	if err != nil {
		return nil, err
	}

	tasks = ct.sort(tasks)

	output := []taskdto.RoutineOutputDTO{}
	for _, task := range tasks {
		if task.TodayRoutine == nil {
			continue
		}

		output = append(output, taskdto.RoutineOutputDTO{
			Id:            task.TodayRoutine.Id,
			ReferenceDate: task.TodayRoutine.ReferenceDate,
			IsFinished:    task.TodayRoutine.IsFinished,
			CreatedAt:     task.TodayRoutine.CreatedAt,
			UpdatedAt:     task.TodayRoutine.UpdatedAt,
			Task: &taskdto.TaskOutputDTO{
				Id:         task.Id,
				Name:       task.Name,
				DueTime:    task.DueTime,
				IsArchived: task.IsArchive,
				CreatedAt:  task.CreatedAt,
				UpdatedAt:  task.UpdatedAt,
			},
		})
	}

	return output, nil

}

func (ct *ListRoutineUseCase) sort(tasks []*entity.Task) []*entity.Task {
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].DueTime < tasks[j].DueTime
	})
	return tasks
}

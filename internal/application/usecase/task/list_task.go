package usecase

import (
	taskdtolist "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/list"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type ListTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewListTaskUseCase(taskRepository repository.ITaskRepository) *ListTaskUseCase {
	return &ListTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *ListTaskUseCase) Execute(isArchived bool) ([]taskdtolist.TaskListOutputDTO, error) {

	tasks, err := ct.TaskRepository.FindAllBy(isArchived)

	if err != nil {
		return nil, err
	}

	output := []taskdtolist.TaskListOutputDTO{}
	for _, task := range tasks {
		output = append(output, taskdtolist.TaskListOutputDTO{
			Id:             task.Id,
			Name:           task.Name,
			DueTime:        task.DueTime,
			IsArchived:     task.IsArchive,
			CompletedTimes: task.CompletedTimes,
			FailedTimes:    task.FailedTimes,
			CreatedAt:      task.CreatedAt,
			UpdatedAt:      task.UpdatedAt,
		})
	}

	return output, nil

}

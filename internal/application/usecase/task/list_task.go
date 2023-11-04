package usecase

import (
	taskdto "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto"
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

func (ct *ListTaskUseCase) Execute(isArchived bool) ([]taskdto.TaskOutputDTO, error) {

	tasks, err := ct.TaskRepository.ListBy(isArchived)

	if err != nil {
		return nil, err
	}

	output := []taskdto.TaskOutputDTO{}
	for _, task := range tasks {
		output = append(output, taskdto.TaskOutputDTO{
			Id:         task.Id,
			Name:       task.Name,
			DueTime:    task.DueTime,
			IsArchived: task.IsArchive,
			CreatedAt:  task.CreatedAt,
			UpdatedAt:  task.UpdatedAt,
		})
	}

	return output, nil

}

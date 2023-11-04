package usecase

import (
	taskdto "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type CreateTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewCreateTaskUseCase(taskRepository repository.ITaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *CreateTaskUseCase) Execute(input *taskdto.TaskInputDTO) (*taskdto.TaskOutputDTO, error) {
	command := &entity.CreateTaskCommand{
		Name:    input.Name,
		DueTime: input.DueTime,
	}

	task, err := entity.CreateTask(command)
	if err != nil {
		return nil, err
	}

	if err := ct.TaskRepository.Create(task); err != nil {
		return nil, err
	}

	output := &taskdto.TaskOutputDTO{
		Id:         task.Id,
		Name:       task.Name,
		DueTime:    task.DueTime,
		IsArchived: task.IsArchive,
		CreatedAt:  task.CreatedAt,
		UpdatedAt:  task.UpdatedAt,
	}

	return output, nil

}

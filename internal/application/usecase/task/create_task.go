package usecase

import (
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type TaskInputDTO struct {
	Name    string `json:"name"`
	DueTime string `json:"due_time"`
}

type TaskOutputDTO struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	DueTime    string    `json:"due_time"`
	IsArchived bool      `json:"is_archived"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewCreateTaskUseCase(taskRepository repository.ITaskRepository) (*CreateTaskUseCase, error) {
	return &CreateTaskUseCase{
		TaskRepository: taskRepository,
	}, nil
}

func (ct *CreateTaskUseCase) Execute(input *TaskInputDTO) (*TaskOutputDTO, error) {
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

	output := &TaskOutputDTO{
		Id:         task.Id,
		Name:       task.Name,
		DueTime:    task.DueTime,
		IsArchived: task.IsArchive,
		CreatedAt:  task.CreatedAt,
		UpdatedAt:  task.UpdatedAt,
	}

	return output, nil

}

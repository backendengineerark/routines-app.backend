package usecase

import "github.com/backendengineerark/routines-app/internal/domain/repository"

type DeleteTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewDeleteTaskUseCase(repository repository.ITaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		TaskRepository: repository,
	}
}

func (at *DeleteTaskUseCase) Execute(taskId string) error {
	task, err := at.TaskRepository.FindById(taskId)
	if err != nil {
		return err
	}

	err = at.TaskRepository.Delete(task)
	if err != nil {
		return err
	}
	return nil

}

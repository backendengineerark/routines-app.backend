package usecase

import "github.com/backendengineerark/routines-app/internal/domain/repository"

type UnarchiveTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewUnarchiveTaskUseCase(repository repository.ITaskRepository) *UnarchiveTaskUseCase {
	return &UnarchiveTaskUseCase{
		TaskRepository: repository,
	}
}

func (at *UnarchiveTaskUseCase) Execute(taskId string) error {
	task, err := at.TaskRepository.FindById(taskId)
	if err != nil {
		return err
	}

	task.Unarchive()

	err = at.TaskRepository.Update(task)
	if err != nil {
		return err
	}
	return nil

}

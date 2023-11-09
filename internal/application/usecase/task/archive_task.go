package usecase

import "github.com/backendengineerark/routines-app/internal/domain/repository"

type ArchiveTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewArchiveTaskUseCase(repository repository.ITaskRepository) *ArchiveTaskUseCase {
	return &ArchiveTaskUseCase{
		TaskRepository: repository,
	}
}

func (at *ArchiveTaskUseCase) Execute(taskId string) error {
	task, err := at.TaskRepository.FindById(taskId)
	if err != nil {
		return err
	}

	task.Archive()

	err = at.TaskRepository.Update(task)
	if err != nil {
		return err
	}
	return nil

}

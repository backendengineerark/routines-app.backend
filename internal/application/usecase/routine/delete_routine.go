package usecase

import (
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type DeleteRoutineUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewDeleteRoutineUseCase(taskRepository repository.ITaskRepository) *DeleteRoutineUseCase {
	return &DeleteRoutineUseCase{
		TaskRepository: taskRepository,
	}
}

func (gr *DeleteRoutineUseCase) Execute(id string) error {

	routine, err := gr.TaskRepository.FindRoutineById(id)
	if err != nil {
		return err
	}

	err = gr.TaskRepository.DeleteRoutine(routine)
	if err != nil {
		return err
	}

	return nil
}

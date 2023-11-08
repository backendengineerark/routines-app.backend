package usecase

import (
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type FinishRoutineUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewFinishRoutineUseCase(taskRepository repository.ITaskRepository) *FinishRoutineUseCase {
	return &FinishRoutineUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *FinishRoutineUseCase) Execute(taskId string) error {

	task, err := ct.TaskRepository.FindById(taskId)
	if err != nil {
		return err
	}

	task.FinishTodayRoutine()

	ct.TaskRepository.UpdateTodayRoutine(task)

	return nil

}

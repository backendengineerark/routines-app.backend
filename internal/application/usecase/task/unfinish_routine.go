package usecase

import (
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type UnfinishRoutineUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewUnfinishRoutineUseCase(taskRepository repository.ITaskRepository) *UnfinishRoutineUseCase {
	return &UnfinishRoutineUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *UnfinishRoutineUseCase) Execute(taskId string) error {

	task, err := ct.TaskRepository.FindById(taskId)
	if err != nil {
		return err
	}

	task.UnfinishTodayRoutine()

	ct.TaskRepository.UpdateTodayRoutine(task)

	return nil

}

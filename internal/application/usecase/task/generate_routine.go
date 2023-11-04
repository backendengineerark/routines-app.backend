package usecase

import (
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type GenerateRoutineUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewGenerateRoutineUseCase(taskRepository repository.ITaskRepository) *GenerateRoutineUseCase {
	return &GenerateRoutineUseCase{
		TaskRepository: taskRepository,
	}
}

func (gr *GenerateRoutineUseCase) Execute() error {

	onlyArchivedTasks := false
	tasks, err := gr.TaskRepository.ListBy(onlyArchivedTasks)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		task.AddRoutine(time.Now())
	}

	return nil

	// command := &entity.CreateTaskCommand{
	// 	Name:    input.Name,
	// 	DueTime: input.DueTime,
	// }

	// task, err := entity.CreateTask(command)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := ct.TaskRepository.Create(task); err != nil {
	// 	return nil, err
	// }

	// output := &taskdto.TaskOutputDTO{
	// 	Id:         task.Id,
	// 	Name:       task.Name,
	// 	DueTime:    task.DueTime,
	// 	IsArchived: task.IsArchive,
	// 	CreatedAt:  task.CreatedAt,
	// 	UpdatedAt:  task.UpdatedAt,
	// }

	// return output, nil

}

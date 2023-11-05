package usecase

import (
	"time"

	taskdto "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type ListRoutineUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewListRoutineUseCase(taskRepository repository.ITaskRepository) *ListRoutineUseCase {
	return &ListRoutineUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *ListRoutineUseCase) Execute(date time.Time) ([]taskdto.RoutineOutputDTO, error) {

	routines, err := ct.TaskRepository.ListRoutine(date)

	if err != nil {
		return nil, err
	}

	output := []taskdto.RoutineOutputDTO{}
	for _, routine := range routines {
		output = append(output, taskdto.RoutineOutputDTO{
			Id:            routine.Id,
			ReferenceDate: routine.ReferenceDate,
			IsFinished:    routine.IsFinished,
			CreatedAt:     routine.CreatedAt,
			UpdatedAt:     routine.UpdatedAt,
			Task: &taskdto.TaskOutputDTO{
				Id:         routine.Task.Id,
				Name:       routine.Task.Name,
				DueTime:    routine.Task.DueTime,
				IsArchived: routine.Task.IsArchive,
				CreatedAt:  routine.Task.CreatedAt,
				UpdatedAt:  routine.Task.UpdatedAt,
			},
		})
	}

	return output, nil

}

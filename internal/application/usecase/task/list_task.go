package usecase

import (
	taskdtolist "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/list"
	weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type ListTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewListTaskUseCase(taskRepository repository.ITaskRepository) *ListTaskUseCase {
	return &ListTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *ListTaskUseCase) Execute(isArchived bool) ([]taskdtolist.TaskListOutputDTO, error) {

	tasks, err := ct.TaskRepository.FindAllBy(isArchived)

	if err != nil {
		return nil, err
	}

	output := []taskdtolist.TaskListOutputDTO{}
	for _, task := range tasks {

		weekdaysOutput := []weekdaydtolist.WeekdayOutputDTO{}
		for _, weekday := range task.Weekdays {
			weekdaysOutput = append(weekdaysOutput, weekdaydtolist.WeekdayOutputDTO{
				Id:        weekday.Id,
				Name:      weekday.Name,
				NumberDay: weekday.NumberDay,
			})
		}

		output = append(output, taskdtolist.TaskListOutputDTO{
			Id:             task.Id,
			Name:           task.Name,
			DueTime:        task.DueTime,
			IsArchived:     task.IsArchive,
			CompletedTimes: task.CompletedTimes,
			FailedTimes:    task.FailedTimes,
			Days:           weekdaysOutput,
			CreatedAt:      task.CreatedAt,
			UpdatedAt:      task.UpdatedAt,
		})
	}

	return output, nil

}

package usecase

import (
	taskdtoupdate "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/update"
	weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type UpdateTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewUpdateTaskUseCase(taskRepository repository.ITaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *UpdateTaskUseCase) Execute(taskId string, input *taskdtoupdate.TaskUpdateInputDTO) (*taskdtoupdate.TaskUpdateOutputDTO, error) {
	task, err := ct.TaskRepository.FindById(taskId)
	if err != nil {
		return nil, err
	}

	var weekdaysIds []string
	for _, weekday := range input.Days {
		weekdaysIds = append(weekdaysIds, weekday.Id)
	}

	weekdays, err := ct.TaskRepository.FindWeekdayIn(weekdaysIds)
	if err != nil {
		return nil, err
	}

	task.ChangeName(input.Name)
	task.ChangeTime(input.DueTime)
	task.NewWeekdays(weekdays)

	if err := ct.TaskRepository.Update(task); err != nil {
		return nil, err
	}

	weekdaysOutput := []weekdaydtolist.WeekdayOutputDTO{}
	for _, weekday := range task.Weekdays {
		weekdaysOutput = append(weekdaysOutput, weekdaydtolist.WeekdayOutputDTO{
			Id:        weekday.Id,
			Name:      weekday.Name,
			NumberDay: weekday.NumberDay,
		})
	}

	output := &taskdtoupdate.TaskUpdateOutputDTO{
		Id:         task.Id,
		Name:       task.Name,
		DueTime:    task.DueTime,
		IsArchived: task.IsArchive,
		Days:       weekdaysOutput,
		CreatedAt:  task.CreatedAt,
		UpdatedAt:  task.UpdatedAt,
	}

	return output, nil

}

package usecase

import (
	taskdtocreate "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/create"
	weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type CreateTaskUseCase struct {
	TaskRepository repository.ITaskRepository
}

func NewCreateTaskUseCase(taskRepository repository.ITaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (ct *CreateTaskUseCase) Execute(input *taskdtocreate.TaskCreateInputDTO) (*taskdtocreate.TaskCreateOutputDTO, error) {
	command := &entity.CreateTaskCommand{
		Name:    input.Name,
		DueTime: input.DueTime,
	}

	var weekdaysIds []string
	for _, weekday := range input.Days {
		weekdaysIds = append(weekdaysIds, weekday.Id)
	}

	weekdays, err := ct.TaskRepository.FindWeekdayIn(weekdaysIds)
	if err != nil {
		return nil, err
	}

	task, err := entity.CreateTask(command, weekdays)
	if err != nil {
		return nil, err
	}

	if err := ct.TaskRepository.Create(task); err != nil {
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

	output := &taskdtocreate.TaskCreateOutputDTO{
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

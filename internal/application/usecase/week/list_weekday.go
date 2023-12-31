package usecase

import (
	weekdaydtolist "github.com/backendengineerark/routines-app/internal/application/usecase/week/dto/list"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type ListWeekdayUseCase struct {
	TaskRepository repository.ITaskRepository
}

func CreateListWeekdayUseCase(repository repository.ITaskRepository) *ListWeekdayUseCase {
	return &ListWeekdayUseCase{
		TaskRepository: repository,
	}
}

func (ld *ListWeekdayUseCase) Execute() ([]weekdaydtolist.WeekdayOutputDTO, error) {
	weekdays, err := ld.TaskRepository.FindAllWeekday()
	if err != nil {
		return nil, err
	}

	output := []weekdaydtolist.WeekdayOutputDTO{}
	for _, weekday := range weekdays {
		output = append(output, weekdaydtolist.WeekdayOutputDTO{
			Id:        weekday.Id,
			Name:      weekday.Name,
			NumberDay: weekday.NumberDay,
		})
	}

	return output, nil
}

package usecase

import (
	"fmt"
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
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

	fmt.Printf("Start generate routine on %s\n", time.Now())

	onlyArchivedTasks := false
	tasks, err := gr.TaskRepository.FindAllBy(onlyArchivedTasks)
	if err != nil {
		return err
	}

	for _, task := range tasks {

		if !gr.shouldGenerateRotineTodayOf(task) {
			fmt.Printf("Today dont have the Task %s\n", task.Name)
			continue
		}

		if task.TodayRoutine != nil {
			fmt.Printf("Task %s already have today routine\n", task.Name)
			continue
		}

		task.CreateTodayRoutine()
		err = gr.TaskRepository.CreateTodayRoutine(task)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("Success to create routine to task %s\n", task.Name)
	}

	fmt.Println("Finished to create routines")

	return nil
}

func (gr *GenerateRoutineUseCase) shouldGenerateRotineTodayOf(task *entity.Task) bool {
	hasTodayRoutine := false
	numberToday := int16(time.Now().Weekday())

	for _, weekday := range task.Weekdays {
		if weekday.NumberDay == numberToday {
			hasTodayRoutine = true
			break
		}
	}

	return hasTodayRoutine
}

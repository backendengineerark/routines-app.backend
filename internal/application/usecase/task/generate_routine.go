package usecase

import (
	"fmt"
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

	fmt.Printf("Start generate routine on %s\n", time.Now())

	onlyArchivedTasks := false
	tasks, err := gr.TaskRepository.FindAllBy(onlyArchivedTasks)
	if err != nil {
		return err
	}

	for _, task := range tasks {

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

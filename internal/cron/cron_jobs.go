package cron

import (
	"fmt"
	"time"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/routine"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/go-co-op/gocron"
)

func ExecuteCronJobs(cron string, taskRepository repository.ITaskRepository) {
	uc := usecase.NewGenerateRoutineUseCase(taskRepository)
	s := gocron.NewScheduler(time.Local)

	s.Cron(cron).Do(func() {
		fmt.Println(time.Now())
		uc.Execute()
	})
	s.StartBlocking()
}

package webhandler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/week"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type WeekdayHandler struct {
	TaskRepository repository.ITaskRepository
}

func NewWeekdayHandler(taskRepository repository.ITaskRepository) *WeekdayHandler {
	return &WeekdayHandler{
		TaskRepository: taskRepository,
	}
}

func (wh *WeekdayHandler) GetWeekday(w http.ResponseWriter, r *http.Request) {

	getWeekdayUseCase := usecase.CreateListWeekdayUseCase(wh.TaskRepository)

	output, err := getWeekdayUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

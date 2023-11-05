package webhandler

import (
	"encoding/json"
	"net/http"
	"time"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/task"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type RoutineHandler struct {
	TaskRepository repository.ITaskRepository
}

func NewRoutineHandler(taskRepository repository.ITaskRepository) *RoutineHandler {
	return &RoutineHandler{
		TaskRepository: taskRepository,
	}
}

func (th *RoutineHandler) ListRoutine(w http.ResponseWriter, r *http.Request) {
	var date time.Time
	if r.URL.Query().Get("date") != "" {
		param := r.URL.Query().Get("date")
		dateCreated, err := time.Parse("2006-01-02", param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		date = dateCreated
	}

	createTaskUseCase := usecase.NewListRoutineUseCase(th.TaskRepository)

	output, err := createTaskUseCase.Execute(date)
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

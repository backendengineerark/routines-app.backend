package webhandler

import (
	"encoding/json"
	"net/http"
	"time"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/routine"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/go-chi/chi/v5"
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

func (th *RoutineHandler) FinishRoutine(w http.ResponseWriter, r *http.Request) {

	taskId := chi.URLParam(r, "task_id")
	if taskId == "" {
		http.Error(w, "task id required", http.StatusBadRequest)
		return
	}

	finishRoutinekUseCase := usecase.NewFinishRoutineUseCase(th.TaskRepository)

	err := finishRoutinekUseCase.Execute(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (th *RoutineHandler) UnfinishRoutine(w http.ResponseWriter, r *http.Request) {

	taskId := chi.URLParam(r, "task_id")
	if taskId == "" {
		http.Error(w, "task id required", http.StatusBadRequest)
		return
	}

	unfinishRoutinekUseCase := usecase.NewUnfinishRoutineUseCase(th.TaskRepository)

	err := unfinishRoutinekUseCase.Execute(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (th *RoutineHandler) DeleteRoutine(w http.ResponseWriter, r *http.Request) {

	taskId := chi.URLParam(r, "routine_id")
	if taskId == "" {
		http.Error(w, "routine id required", http.StatusBadRequest)
		return
	}

	deleteRoutinekUseCase := usecase.NewDeleteRoutineUseCase(th.TaskRepository)

	err := deleteRoutinekUseCase.Execute(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

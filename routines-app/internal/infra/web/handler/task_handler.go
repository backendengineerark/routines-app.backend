package webhandler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/task"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type TaskHandler struct {
	TaskRepository repository.ITaskRepository
}

func NewTaskHandler(taskRepository repository.ITaskRepository) *TaskHandler {
	return &TaskHandler{
		TaskRepository: taskRepository,
	}
}

func (th *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.TaskInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	createTaskUseCase, err := usecase.NewCreateTaskUseCase(th.TaskRepository)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := createTaskUseCase.Execute(&input)
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

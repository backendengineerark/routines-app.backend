package webhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/task"
	taskdto "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto"
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
	var input taskdto.TaskInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	createTaskUseCase := usecase.NewCreateTaskUseCase(th.TaskRepository)
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

func (th *TaskHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	isArchivedParam := false

	if r.URL.Query().Get("is_archived") != "" {
		param, err := strconv.ParseBool(r.URL.Query().Get("is_archived"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		isArchivedParam = param
	}

	createTaskUseCase := usecase.NewListTaskUseCase(th.TaskRepository)

	output, err := createTaskUseCase.Execute(isArchivedParam)
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

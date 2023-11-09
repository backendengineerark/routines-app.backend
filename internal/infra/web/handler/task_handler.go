package webhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/task"
	taskdtocreate "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/create"
	taskdtoupdate "github.com/backendengineerark/routines-app/internal/application/usecase/task/dto/update"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
	"github.com/go-chi/chi/v5"
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
	var input taskdtocreate.TaskCreateInputDTO
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

func (th *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "task_id")

	var input taskdtoupdate.TaskUpdateInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	updateTaskUseCase := usecase.NewUpdateTaskUseCase(th.TaskRepository)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := updateTaskUseCase.Execute(taskIdParam, &input)
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

func (th *TaskHandler) Archive(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "task_id")

	createTaskUseCase := usecase.NewArchiveTaskUseCase(th.TaskRepository)

	err := createTaskUseCase.Execute(taskIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (th *TaskHandler) Unarchive(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "task_id")

	createTaskUseCase := usecase.NewUnarchiveTaskUseCase(th.TaskRepository)

	err := createTaskUseCase.Execute(taskIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (th *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "task_id")

	createTaskUseCase := usecase.NewDeleteTaskUseCase(th.TaskRepository)

	err := createTaskUseCase.Execute(taskIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

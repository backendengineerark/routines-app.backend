package webhandler

import (
	"encoding/json"
	"net/http"
	"time"

	usecase "github.com/backendengineerark/routines-app/internal/application/usecase/metric"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type MetricHandler struct {
	MetricRepository repository.IMetricRepository
}

func NewMetricHandler(taskRepository repository.IMetricRepository) *MetricHandler {
	return &MetricHandler{
		MetricRepository: taskRepository,
	}
}

func (th *MetricHandler) GetMetric(w http.ResponseWriter, r *http.Request) {
	var initial time.Time
	var end time.Time

	if r.URL.Query().Get("initial_date") != "" {
		param := r.URL.Query().Get("initial_date")
		dateCreated, err := time.Parse("2006-01-02", param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		initial = dateCreated
	}

	if r.URL.Query().Get("end_date") != "" {
		param := r.URL.Query().Get("end_date")
		dateCreated, err := time.Parse("2006-01-02", param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		end = dateCreated
	}

	getMetricUseCase := usecase.NewGetMetricUseCase(th.MetricRepository)

	output, err := getMetricUseCase.Execute(initial, end)
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

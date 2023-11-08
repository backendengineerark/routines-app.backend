package usecase

import (
	"fmt"
	"time"

	metricdto "github.com/backendengineerark/routines-app/internal/application/usecase/metric/dto"
	"github.com/backendengineerark/routines-app/internal/domain/repository"
)

type GetMetricUseCase struct {
	MetricRepository repository.IMetricRepository
}

func NewGetMetricUseCase(metricRepository repository.IMetricRepository) *GetMetricUseCase {
	return &GetMetricUseCase{
		MetricRepository: metricRepository,
	}
}

func (ct *GetMetricUseCase) Execute(initial time.Time, end time.Time) ([]metricdto.MetricOutputDTO, error) {

	output := []metricdto.MetricOutputDTO{}

	metrics, err := ct.MetricRepository.FindAllRoutinesByRangeDate(initial, end)

	if err != nil {
		return nil, err
	}

	days := getListDays(initial, end)

	for _, day := range days {
		hasRoutine := false
		metricsOutput := []metricdto.MetricDTO{}
		for _, metric := range metrics {

			if day.Equal(metric.ReferenceDate) {
				hasRoutine = true
				metricsOutput = append(metricsOutput, metricdto.MetricDTO{
					Id:         metric.Id,
					TaskName:   metric.TaskName,
					IsFinished: metric.IsFinished,
				})
			}
		}

		if hasRoutine {
			output = append(output, metricdto.MetricOutputDTO{
				Date:     day.Format("2006-01-02"),
				Routines: metricsOutput,
			})
		}
	}

	fmt.Println(days)

	for _, metric := range metrics {
		fmt.Printf(metric.Id)
	}

	return output, nil

}

func getListDays(initial time.Time, end time.Time) []time.Time {
	days := end.Sub(initial).Hours() / 24
	var dates []time.Time

	for i := 0; i <= int(days); i++ {
		dates = append(dates, initial.AddDate(0, 0, i))
	}
	return dates
}

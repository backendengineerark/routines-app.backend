package database

import (
	"database/sql"
	"time"

	"github.com/backendengineerark/routines-app/internal/domain/metric/entity"
)

type MetricMysqlRepository struct {
	DB *sql.DB
}

func NewMetricMysqlRepository(db *sql.DB) *MetricMysqlRepository {
	return &MetricMysqlRepository{
		DB: db,
	}
}

func (tr *MetricMysqlRepository) FindAllRoutinesByRangeDate(initial time.Time, end time.Time) ([]*entity.Metric, error) {
	stmt, err := tr.DB.Prepare("select r.id, t.name, r.is_finished, r.reference_date from routines r inner join tasks t on t.id  = r.tasks_id where r.reference_date between ? and ? order by r.reference_date asc")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(initial, end)
	if err != nil {
		return nil, err
	}

	metrics := []*entity.Metric{}

	for rows.Next() {
		var metric entity.Metric

		err := rows.Scan(&metric.Id, &metric.TaskName, &metric.IsFinished, &metric.ReferenceDate)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, &metric)
	}

	return metrics, nil
}

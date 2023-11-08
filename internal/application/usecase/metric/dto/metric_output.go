package metricdto

type MetricDTO struct {
	Id         string `json:"id"`
	TaskName   string `json:"task_name"`
	IsFinished bool   `json:"is_finished"`
}

type MetricOutputDTO struct {
	Date     string      `json:"date"`
	Routines []MetricDTO `json:"routines"`
}

package entity

import "time"

type Metric struct {
	Id            string
	TaskName      string
	IsFinished    bool
	ReferenceDate time.Time
}

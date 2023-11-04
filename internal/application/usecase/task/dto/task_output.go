package taskdto

import "time"

type TaskOutputDTO struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	DueTime    string    `json:"due_time"`
	IsArchived bool      `json:"is_archived"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

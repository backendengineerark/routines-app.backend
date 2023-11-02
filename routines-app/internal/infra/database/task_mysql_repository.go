package database

import (
	"database/sql"

	"github.com/backendengineerark/routines-app/internal/domain/task/entity"
)

type TaskMysqlRepository struct {
	DB *sql.DB
}

func NewTaskMysqlRepository(db *sql.DB) *TaskMysqlRepository {
	return &TaskMysqlRepository{
		DB: db,
	}
}

func (tr *TaskMysqlRepository) Create(task *entity.Task) error {
	stmt, err := tr.DB.Prepare("insert into tasks(id, name, due_time, is_archived, created_at, updated_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Id, task.Name, task.DueTime, task.IsArchive, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskMysqlRepository) Update(task *entity.Task) error {
	stmt, err := tr.DB.Prepare("update tasks set name = ?, due_time = ?, is_archived = ?, updated_at = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Name, task.DueTime, task.IsArchive, task.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// func (tr *TaskMysqlRepository) listsBy(isArchived bool) ([]entity.Task, error) {
// 	rows, err := tr.DB.Query("select id, name, due_time, is_archived, created_at, updated_at where is_archived = ?")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var tasks []entity.Task

// 	for rows.Next() {
// 		var task entity.Task

// 		err := rows.Scan(&task.Id, &task.Name, &task.DueTime, &task.IsArchive, &task.CreatedAt, &task.UpdatedAt)

// 		if err != nil {
// 			return nil, err
// 		}

// 		tasks = append(tasks, task)
// 	}

// 	return tasks, nil
// }

package database

import (
	"database/sql"
	"time"

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

	stmt, err = tr.DB.Prepare("insert into routines(id, tasks_id, reference_date, is_finished, created_at, updated_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Routines[0].Id, task.Id, task.Routines[0].ReferenceDate, task.Routines[0].IsFinished, task.Routines[0].CreatedAt, task.Routines[0].Task.UpdatedAt)
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

func (tr *TaskMysqlRepository) ListBy(isArchived bool) ([]entity.Task, error) {
	stmt, err := tr.DB.Prepare("select id, name, due_time, is_archived, created_at, updated_at from tasks where is_archived = ? order by created_at desc")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(isArchived)
	if err != nil {
		return nil, err
	}

	tasks := []entity.Task{}

	for rows.Next() {
		var task entity.Task

		err := rows.Scan(&task.Id, &task.Name, &task.DueTime, &task.IsArchive, &task.CreatedAt, &task.UpdatedAt)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskMysqlRepository) ListRoutine(date time.Time) ([]entity.Routine, error) {
	stmt, err := tr.DB.Prepare("SELECT r.id, r.reference_date, r.is_finished, r.created_at, r.updated_at, t.id, t.name, t.due_time, t.is_archived, t.created_at, t.updated_at from routines r inner join tasks t on t.id = r.tasks_id where reference_date = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(date)
	if err != nil {
		return nil, err
	}

	routines := []entity.Routine{}

	for rows.Next() {
		var routine entity.Routine
		var task entity.Task

		err := rows.Scan(&routine.Id, &routine.ReferenceDate, &routine.IsFinished, &routine.CreatedAt, &routine.UpdatedAt, &task.Id, &task.Name, &task.DueTime, &task.IsArchive, &task.CreatedAt, &task.UpdatedAt)

		if err != nil {
			return nil, err
		}

		routine.Task = &task

		routines = append(routines, routine)
	}

	return routines, nil
}

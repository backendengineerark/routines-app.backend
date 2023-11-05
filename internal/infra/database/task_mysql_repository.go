package database

import (
	"database/sql"

	"github.com/backendengineerark/routines-app/internal/domain/common/custom_dates"
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

	_, err = stmt.Exec(task.TodayRoutine.Id, task.Id, task.TodayRoutine.ReferenceDate, task.TodayRoutine.IsFinished, task.TodayRoutine.CreatedAt, task.TodayRoutine.UpdatedAt)
	if err != nil {
		return err
	}

	tr.CreateTodayRoutine(task)

	return nil
}

func (tr *TaskMysqlRepository) Update(task *entity.Task) error {
	stmt, err := tr.DB.Prepare("update tasks set name = ?, due_time = ?, is_archived = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Name, task.DueTime, task.IsArchive, task.UpdatedAt, task.Id)
	if err != nil {
		return err
	}

	var exists = ""
	err = tr.DB.QueryRow("select id from routines where id = ?").Scan(exists)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskMysqlRepository) CreateTodayRoutine(task *entity.Task) error {
	stmt, err := tr.DB.Prepare("insert into routines(id, tasks_id, reference_date, is_finished, created_at, updated_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.TodayRoutine.Id, task.Id, task.TodayRoutine.ReferenceDate, task.TodayRoutine.IsFinished, task.TodayRoutine.CreatedAt, task.TodayRoutine.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskMysqlRepository) UpdateTodayRoutine(task *entity.Task) error {
	stmt, err := tr.DB.Prepare("update routines set is_finished = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.TodayRoutine.IsFinished, task.TodayRoutine.UpdatedAt, task.TodayRoutine.Id)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskMysqlRepository) FindById(id string) (*entity.Task, error) {
	stmt, err := tr.DB.Prepare("SELECT t.id, t.name, t.due_time, t.is_archived, t.created_at, t.created_at, from tasks t where t.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var task entity.Task

	err = stmt.QueryRow(id).Scan(&task.Id, &task.Name, &task.DueTime, &task.IsArchive, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	routine, err := tr.FindTodayRoutineByTask(task)
	if err == nil {
		task.TodayRoutine = routine
	}

	return &task, nil
}

func (tr *TaskMysqlRepository) FindAllBy(isArchived bool) ([]*entity.Task, error) {
	stmt, err := tr.DB.Prepare("SELECT t.id, t.name, t.due_time, t.is_archived, t.created_at, t.updated_at from tasks t where t.is_archived = ? order by t.created_at desc")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(isArchived)
	if err != nil {
		return nil, err
	}

	tasks := []*entity.Task{}

	for rows.Next() {
		var task entity.Task

		err := rows.Scan(&task.Id, &task.Name, &task.DueTime, &task.IsArchive, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		routine, err := tr.FindTodayRoutineByTask(task)
		if err == nil {
			task.TodayRoutine = routine
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (tr *TaskMysqlRepository) FindTodayRoutineByTask(task entity.Task) (*entity.Routine, error) {
	var routine entity.Routine
	today := custom_dates.TodayBeginningHour()

	stmt, err := tr.DB.Prepare("SELECT r.id, r.reference_date, r.is_finished, r.created_at, r.updated_at from routines r where r.tasks_id = ? and r.reference_date = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(task.Id, today).Scan(&routine.Id, &routine.ReferenceDate, &routine.IsFinished, &routine.CreatedAt, &routine.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &routine, nil
}

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

	return nil
}

func (tr *TaskMysqlRepository) Delete(task *entity.Task) error {
	err := tr.deleteAllroutinesBy(task)
	if err != nil {
		return err
	}

	stmt, err := tr.DB.Prepare("delete from tasks where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Id)
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
	stmt, err := tr.DB.Prepare("select t.id, t.name, t.due_time, t.is_archived, t.created_at, t.created_at from tasks t where t.id = ?")
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
	stmt, err := tr.DB.Prepare("select t.id, t.name, t.due_time, t.is_archived, t.created_at, t.updated_at, (SELECT COUNT(*) FROM routines r where is_finished = true and r.tasks_id  = t.id) as completed_times, (SELECT COUNT(*) FROM routines r where is_finished = false and r.tasks_id  = t.id) as failed_times from tasks t where t.is_archived = ? order by completed_times asc")
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

		err := rows.Scan(&task.Id, &task.Name, &task.DueTime, &task.IsArchive, &task.CreatedAt, &task.UpdatedAt, &task.CompletedTimes, &task.FailedTimes)
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

func (tr *TaskMysqlRepository) FindAllWeekday() ([]*entity.Weekday, error) {
	rows, err := tr.DB.Query("select id, name, number_day from weekdays order by number_day asc")
	if err != nil {
		return nil, err
	}

	var weekdays []*entity.Weekday

	for rows.Next() {
		var weekday entity.Weekday

		err := rows.Scan(&weekday.Id, &weekday.Name, &weekday.NumberDay)
		if err != nil {
			return nil, err
		}
		weekdays = append(weekdays, &weekday)
	}

	return weekdays, nil
}

func (tr *TaskMysqlRepository) FindTodayRoutineByTask(task entity.Task) (*entity.Routine, error) {
	var routine entity.Routine
	today := custom_dates.TodayBeginningHour()

	stmt, err := tr.DB.Prepare("select r.id, r.reference_date, r.is_finished, r.created_at, r.updated_at from routines r where r.tasks_id = ? and r.reference_date = ?")
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

func (tr *TaskMysqlRepository) deleteAllroutinesBy(task *entity.Task) error {
	stmt, err := tr.DB.Prepare("delete from routines where tasks_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Id)
	if err != nil {
		return err
	}

	return nil
}

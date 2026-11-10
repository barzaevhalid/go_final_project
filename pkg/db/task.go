package db

import (
	"database/sql"
	"errors"
	"fmt"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4)`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err

}

func Tasks(limit int) ([]*Task, error) {

	rows, err := DB.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ? ", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if tasks == nil {
		return make([]*Task, 0), nil
	}
	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	var task Task
	err := DB.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id=$1", id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *Task) error {

	query := `UPDATE scheduler SET  title = $1, comment = $2, date = $3, repeat = $4
			WHERE id=$5`
	res, err := DB.Exec(query, task.Title, task.Comment, task.Date, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func DoneTask(id string) error {
	task, err := GetTask(id)

	if err != nil {
		return err
	}

	if task.Repeat == "" {
		err := DeleteTask(id)
		return err

	}
	return err
}

func DeleteTask(id string) error {
	result, err := DB.Exec("DELETE FROM scheduler WHERE id = $1", id)
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func UpdateDate(next string, id string) error {

	_, err := DB.Exec("UPDATE scheduler SET date = $1 WHERE id = $2", next, id)
	return err
}

package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db pgxpool.Pool

type Task struct {
	id          int
	opened      int
	closed      int
	author_id   int
	assigned_id int
	title       string
}

func NewTask(title string) (int, error) {
	var id int
	err := db.QueryRow(context.Background(), `INSERT INTO tasks (name, title) 
	VALUES ($1) RETURNING id;`, title).Scan(&id)
	if err != nil {
		fmt.Println("could not insert new task")
	}

	return id, err
}

var tasks []Task

func AllTasks(db *pgxpool.Pool) ([]Task, error) {
	rows, err := db.Query(context.Background(), `select * from task;`)
	if err != nil {
		fmt.Println("Query error")
	}
	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.id,
			&task.opened,
			&task.closed,
			&task.author_id,
			&task.assigned_id,
			&task.title,
		)
		if err != nil {
			fmt.Println("Scan error")
		}
		tasks = append(tasks, task)
	}
	return tasks, rows.Err()
}

func TaskOnAuthor(name string) ([]Task, error) {

	rows, err := db.Query(context.Background(), `SELECT * FROM tasks WHERE author_id IN 
	(SELECT id FROM users WHERE  name = $1);`, name)
	if err != nil {
		fmt.Println("could not select tasks on author")
	}

	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.id,
			&task.opened,
			&task.closed,
			&task.author_id,
			&task.assigned_id,
			&task.title,
		)
		if err != nil {
			fmt.Println("Scan error")
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func TaskOnLable(name string) ([]Task, error) {

	rows, err := db.Query(context.Background(), `SELECT * FROM tasks WHERE id IN 
	(SELECT tasks_lables.task_id FROM tasks_lables WHERE tasks_lables.lables_id IN  
	(SELECT labels.id FROM labels WHERE  name = $1));`, name)
	if err != nil {
		fmt.Println("could not select tasks on author")
	}

	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.id,
			&task.opened,
			&task.closed,
			&task.author_id,
			&task.assigned_id,
			&task.title,
		)
		if err != nil {
			fmt.Println("Scan error")
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func UpdateTask(id int, title string) error {

	_, err := db.Exec(context.Background(), `UPDATE tasks SET title = $2 
	WHERE id = $1;`, id, title)
	if err != nil {
		fmt.Println("could not update task")
	}
	return nil
}

func DeleteTask(id int) error {
	_, err := db.Exec(context.Background(), `DELETE FROM task WHERE id = $1;`, id)
	if err != nil {
		fmt.Println("could not delete task")
	}

	return nil
}

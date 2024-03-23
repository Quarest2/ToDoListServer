package pg_storage

import "time"

func CreateTask(task Task) (Task, error) {
	sqlStatement := `INSERT INTO tasks (username, task, deadline, is_done)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, task.Username, task.Task, task.Deadline, task.Is_done).Scan(&id)
	if err != nil {
		return Task{}, err
	}
	task.Id = id
	return task, nil
}

func EditTask(task Task, id int) (Task, error) {
	sqlStatement := `UPDATE tasks
		SET username = $2, task = $3, deadline = $4, is_done = $5
		WHERE id = $1
		RETURNING id`
	err := db.QueryRow(sqlStatement, id, task.Username, task.Task, task.Deadline, task.Is_done).Scan(&id)
	if err != nil {
		return Task{}, err
	}
	task.Id = id
	return task, nil
}

func DeleteTask(id int) (Task, error) {
	sqlStatement := `DELETE FROM tasks 
       WHERE id = $1
       RETURNING id, username, task, deadline, is_done`
	delID := 0
	delUsername := ""
	delTask := ""
	delDeadline := time.Time{}
	delIsDone := false
	err := db.QueryRow(sqlStatement, id).Scan(&delID, &delUsername, &delTask, &delDeadline, &delIsDone)
	if err != nil {
		return Task{}, err
	}
	task := Task{Id: delID, Username: delUsername, Task: delTask, Deadline: delDeadline, Is_done: delIsDone}
	return task, nil
}

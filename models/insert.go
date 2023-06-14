package models

import (
	"github.com/zangassis/go-todo-api/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		println(err)
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURN id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}

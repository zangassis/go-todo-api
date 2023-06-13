package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zangassis/go-todo-api/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprint("host=%s port=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}

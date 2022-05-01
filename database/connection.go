package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() sql.DB {
	connection, err := sql.Open("mysql", "yt_go_auth:123456@/yt_go_auth")
	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	// See "Important settings" section.
	connection.SetConnMaxLifetime(time.Minute * 3)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)
	return *connection
}

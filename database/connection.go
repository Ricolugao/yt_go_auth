package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	connection, err := sql.Open("mysql", "root:123456@/yt_go_auth")
	if err != nil {
		panic(err)
	}
	fmt.Println(connection)
	// See "Important settings" section.
	connection.SetConnMaxLifetime(time.Minute * 3)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)
}

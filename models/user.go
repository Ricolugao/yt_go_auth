package models

import (
	"fmt"
	"log"
	"yt_go_auth/database"
)

type User struct {
	Id       uint
	Name     string
	Email    string
	Password []byte
}

func InsereNovoUsuario(user User) User {
	var userInserido User
	userInserido.Name = user.Name
	userInserido.Email = user.Email
	userInserido.Password = user.Password

	db := database.Connect()
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(userInserido.Name, userInserido.Email, userInserido.Password)
	if err != nil {
		log.Println(err)
	}

	lastId, _ := result.LastInsertId()

	userInserido.Id = uint(lastId)

	return userInserido
}

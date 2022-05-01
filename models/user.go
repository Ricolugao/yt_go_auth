package models

import (
	"fmt"
	"log"
	"yt_go_auth/database"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
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

func Logar(data User) User {
	var usuarioLogado User

	db := database.Connect()
	defer db.Close()

	queryString := "SELECT * FROM users where email = ?"

	err := db.QueryRow(queryString, data.Email).Scan(&usuarioLogado.Id, &usuarioLogado.Name, &usuarioLogado.Email, &usuarioLogado.Password)
	if err != nil {
		log.Println(err)
	}

	return usuarioLogado
}

func FindUserById(id string) User {
	var usuarioLogado User

	db := database.Connect()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM users where id = ?", id).Scan(&usuarioLogado.Id, &usuarioLogado.Name, &usuarioLogado.Email, &usuarioLogado.Password)
	if err != nil {
		log.Println(err)
	}
	return usuarioLogado
}

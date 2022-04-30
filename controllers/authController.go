package controllers

import (
	"yt_go_auth/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	novoUsuario := models.InsereNovoUsuario(user)

	return c.JSON(novoUsuario)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	informacoesDeLogin := models.User{
		Email:    data["email"],
		Password: password,
	}

	usuarioLogado := models.Logar(informacoesDeLogin)

	if usuarioLogado.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Usuário não encontrado",
		})
	}

	if err := bcrypt.CompareHashAndPassword(usuarioLogado.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Senha incorreta",
		})
	}

	return c.JSON(usuarioLogado)
}

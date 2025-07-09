package handler

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
)

func Adduser(a *fiber.Ctx) error {
	var user model.User

	err := a.BodyParser(&user)
	if err != nil {
		return err
	}
	database.Db.Create(&user)
	return a.SendStatus(200)
}

func Getuser(a *fiber.Ctx) error {
	var users []model.User

	database.Db.Find(&users)
	return a.JSON(users)
}

func Deleteuser(a *fiber.Ctx) error {

	var user model.User
	a.BodyParser(&user)

	var id = user.ID

	database.Db.Delete(&user, id)
	return a.SendStatus(200)
}

func Updateuser(a *fiber.Ctx) error {

	var user model.User
	a.BodyParser(&user)
	var id = user.ID

	database.Db.Where("id = ?", id).Updates(&user)
	return a.SendStatus(200)
}

package handler

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
)

func Countemp(a *fiber.Ctx) error {
	var count int64
	err := database.Db.Table("employees").Count(&count).Error
	if err != nil {
		return err
	}
	return a.JSON(fiber.Map{"count": count})
}
func Adduser(a *fiber.Ctx) error {
	var user model.User

	err := a.BodyParser(&user)
	if err != nil {
		return err
	}
	err = database.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return a.SendStatus(201)
}

func Getusers(a *fiber.Ctx) error {
	var users []model.User

	database.Db.Find(&users)
	return a.JSON(users)
}

func Deluser(a *fiber.Ctx) error {

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

func Addemp(a *fiber.Ctx) error {
	var emp model.Employee

	err := a.BodyParser(&emp)
	if err != nil {
		return err
	}
	err = database.Db.Create(&emp).Error
	if err != nil {
		return err
	}
	return a.SendStatus(200)
}

func Getemps(a *fiber.Ctx) error {
	var emps []model.Employee

	database.Db.Find(&emps)
	return a.JSON(emps)
}

func Delemp(a *fiber.Ctx) error {
	var emp model.Employee
	a.BodyParser(&emp)

	var id = emp.ID

	database.Db.Delete(&emp, id)
	return a.SendStatus(200)
}

func Addadm(a *fiber.Ctx) error {
	var adm model.Admin

	err := a.BodyParser(&adm)
	if err != nil {
		return err
	}
	err = database.Db.Create(&adm).Error
	if err != nil {
		return err
	}
	return a.SendStatus(200)
}

func Getadm(a *fiber.Ctx) error {
	var adm []model.Admin

	database.Db.Find(&adm)
	return a.JSON(adm)
}

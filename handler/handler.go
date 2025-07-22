package handler

import (
	"api/database"
	"api/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Countemp(c *fiber.Ctx) error {
	var count int64
	err := database.Db.Table("employees").Count(&count).Error
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"count": count})
}
func Adduser(c *fiber.Ctx) error {
	var user model.User

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing password")
	}

	user.Password = string(hashedPassword)

	err = database.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return c.SendStatus(200)
}

func Getusers(c *fiber.Ctx) error {
	var users []model.User

	database.Db.Find(&users)
	return c.JSON(users)
}

func Deluser(c *fiber.Ctx) error {

	var user model.User
	c.BodyParser(&user)

	var id = user.ID

	database.Db.Delete(&user, id)
	return c.SendStatus(200)
}

func Updateuser(c *fiber.Ctx) error {

	var user model.User
	c.BodyParser(&user)
	var id = user.ID

	database.Db.Where("id = ?", id).Updates(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func Addemp(c *fiber.Ctx) error {
	var emp model.Employee

	err := c.BodyParser(&emp)
	if err != nil {
		return err
	}
	err = database.Db.Create(&emp).Error
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(emp)
}

func Getemps(c *fiber.Ctx) error {
	var emps []model.Employee

	database.Db.Find(&emps)
	return c.Status(fiber.StatusOK).JSON(emps)
}

func Delemp(c *fiber.Ctx) error {
	var emp model.Employee
	c.BodyParser(&emp)

	var id = emp.ID

	database.Db.Delete(&emp, id)
	return c.Status(fiber.StatusOK).JSON(emp)
}

func Updemp(c *fiber.Ctx) error {

	var emp model.Employee
	c.BodyParser(&emp)
	var id = emp.ID

	database.Db.Where("id = ?", id).Updates(&emp)
	return c.Status(fiber.StatusOK).JSON(emp)
}

func Addadm(c *fiber.Ctx) error {
	var adm model.Admin

	err := c.BodyParser(&adm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err = database.Db.Create(&adm).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create admin",
		})
	}
	return c.Status(fiber.StatusOK).JSON(adm)
}

func Getadm(c *fiber.Ctx) error {
	var adm []model.Admin

	database.Db.Find(&adm)
	return c.Status(fiber.StatusOK).JSON(adm)
}

func Login(c *fiber.Ctx) error {
	var input model.Log
	var user model.User

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	result := database.Db.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid creds")
	}
	fmt.Print(err)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login succesful",
	})
}

func Adddep(c *fiber.Ctx) error {
	var dep model.Dept
	err := c.BodyParser(&dep)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err = database.Db.Create(&dep).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create department",
		})
	}
	return c.Status(fiber.StatusOK).JSON(dep)
}

func Getdep(c *fiber.Ctx) error {
	var dep []model.Dept

	database.Db.Find(&dep)
	return c.Status(fiber.StatusOK).JSON(dep)
}

func Countdepemp(c *fiber.Ctx) error {
	department := c.Query("department")
	var count int64
	if department == "" {
		database.Db.Model(&model.Employee{}).Count(&count)
	} else {
		database.Db.Model(&model.Employee{}).Where("department = ?", department).Count(&count)
	}
	return c.JSON(fiber.Map{"count": count})
}

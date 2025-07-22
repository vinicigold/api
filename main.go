package main

import (
	"api/database"
	"api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	fiber := fiber.New()
	fiber.Use(cors.New())
	fiber.Post("/Login", handler.Login)
	fiber.Get("/countemp", handler.Countemp)
	fiber.Post("/adduser", handler.Adduser)
	fiber.Get("/getusers", handler.Getusers)
	fiber.Delete("/deluser", handler.Deluser)
	fiber.Put("/upuser", handler.Updateuser)
	fiber.Post("/addemp", handler.Addemp)
	fiber.Get("/getemps", handler.Getemps)
	fiber.Put("/upemp", handler.Updemp)
	fiber.Delete("/delemp", handler.Delemp)
	fiber.Post("/addadmin", handler.Addadm)
	fiber.Get("/getadmin", handler.Getadm)
	fiber.Post("/adddep", handler.Adddep)
	fiber.Get("/getdep", handler.Getdep)
	fiber.Get("/countdepemp", handler.Countdepemp)

	fiber.Listen(":3002")
}

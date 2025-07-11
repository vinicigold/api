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
	fiber.Get("/countuser", handler.Countuser)
	fiber.Post("/adduser", handler.Adduser)
	fiber.Get("/getusers", handler.Getuser)
	fiber.Delete("/deluser", handler.Deleteuser)
	fiber.Put("/upuser", handler.Updateuser)
	fiber.Post("/addemp", handler.Addemp)
	fiber.Get("/getemps", handler.Getemps)
	fiber.Post("/addadmin", handler.Addadm)
	fiber.Get("/getadmin", handler.Getadm)

	fiber.Listen(":3002")
}

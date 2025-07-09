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
	fiber.Post("/user", handler.Adduser)
	fiber.Get("/users", handler.Getuser)
	fiber.Delete("/user", handler.Deleteuser)
	fiber.Put("/up", handler.Updateuser)

	fiber.Listen(":3002")
}

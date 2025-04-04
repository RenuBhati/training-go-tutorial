package main

import (
	"fmt"

	"github.com/RenuBhati/training-go-tutorial/database"
	"github.com/RenuBhati/training-go-tutorial/models"
	"github.com/RenuBhati/training-go-tutorial/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	fmt.Println("Program Started")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layout",
	})

	database.Connect()

	models.Migrate(database.DB)

	routes.SetupRoutes(app)

	app.Use(logger.New())
	app.Use(recover.New())

	app.Listen(":8080")
	fmt.Println("Server Started")
}

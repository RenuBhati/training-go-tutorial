package routes

import (
	"fmt"

	"github.com/RenuBhati/training-go-tutorial/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	fmt.Println("Routes")
	app.Get("/", controllers.GetHome)
	app.Get("/users", controllers.GetUsers)
	app.Get("/api/users", controllers.GetUsersJSON)  

}

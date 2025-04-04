package controllers

import (
	"github.com/RenuBhati/training-go-tutorial/database"
	"github.com/RenuBhati/training-go-tutorial/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsersJSON(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.Render("view_user", fiber.Map{
		"Users": users,
	})
}

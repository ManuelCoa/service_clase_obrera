package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostLogin(c *fiber.Ctx) error {
	return handlers.HandleLogin(c, models.SessionLogin)
}

func PostLogout(c *fiber.Ctx) error {
	return handlers.HandleLogout(c, models.SessionLogout)
}

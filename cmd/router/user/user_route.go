package router

import (
	handler "king/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userhandler handler.UserHandler) {
	app.Post("/register", userhandler.RegisterUser)
	app.Get("/user/:username", userhandler.GetUserName)
}
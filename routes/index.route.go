package routes

import "github.com/gofiber/fiber/v2"

func InitRoute(app *fiber.App) {	
	v1Route(app)
}
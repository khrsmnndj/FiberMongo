package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/khrsmnndj/FiberMongo/api/oauth/service"
)

func SetupRoutes(app fiber.Router) {
	oauth := app.Group("/v1/oauth")
	oauth.Post("/authenticate", service.Authenticate)
}

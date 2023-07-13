package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrsmnndj/FiberMongo/api/oauth/query"
)

func Authenticate(c *fiber.Ctx) error {
	users := query.FindAllUser(c)
	return c.JSON(fiber.Map{"data": users})
}
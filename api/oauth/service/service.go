package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrsmnndj/FiberMongo/api/oauth/repository"
)

// Hello handle api status
func Authenticate(c *fiber.Ctx) error {
	return repository.Authenticate(c)
}

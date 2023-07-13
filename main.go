package main

import (

    "github.com/gofiber/fiber/v2"

    "github.com/khrsmnndj/FiberMongo/infra/config"

    // router
    oauth "github.com/khrsmnndj/FiberMongo/api/oauth/router"
)

func main() {

    app := fiber.New()
    api := app.Group("/api")

    app.Static("/file", "./public")
    
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "success", "message": "Server is running properly!", "data": nil})
    })

    // init route
    oauth.SetupRoutes(api)

    app.Listen(":" + config.PORT)
}
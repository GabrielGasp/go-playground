package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func handler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	errorMsg := "Something went wrong!"

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code == fiber.StatusNotFound {
		errorMsg = "Not found"
	}

	// Send error page
	return c.Status(code).JSON(fiber.Map{
		"error": errorMsg,
	})
}

func fiberSetup() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// recover middleware
	app.Use(recover.New())
	// logger middleware
	app.Use(logger.New())

	app.Get("/", handler)

	return app
}

func main() {
	app := fiberSetup()

	app.Listen(":3000")
}

package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Teste struct {
	A int `json:"a,omitempty"`
	B int `json:"b,omitempty"`
	C int `json:"c,omitempty"`
}

func rng(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1
}

func handler(c *fiber.Ctx) error {
	// initialize struct
	t := new(Teste)

	// get query params
	incA := c.Query("a")
	incB := c.Query("b")
	incC := c.Query("c")

	// insert values
	if incA == "true" {
		t.A = rng(100)
	}
	if incB == "true" {
		t.B = rng(100)
	}
	if incC == "true" {
		t.C = rng(100)
	}

	return c.JSON(t)
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

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// recover middleware
	app.Use(recover.New())
	// logger middleware
	app.Use(logger.New())

	app.Get("/teste", handler)

	app.Listen(":3000")
}

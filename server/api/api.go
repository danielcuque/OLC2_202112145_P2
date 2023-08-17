package api

import (
	I "OLC2/chore/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Response will have symbols, errors, logs and cst
type Resp struct {
	Symbols []I.ApiVariable   `json:"symbols"`
	Errors  []*I.VisitorError `json:"errors"`
	Logs    []string          `json:"logs"`
	Cst     string            `json:"cst"`
}

type Message struct {
	Content string `json:"content"`
}

// Return errors as json

func HandleVisitor(c *fiber.Ctx) error {

	code := c.FormValue("code")

	result := I.NewEvaluator(code)

	response := Resp{
		Symbols: result.Env.GetSymbolTable(),
		Errors:  result.Errors,
		Logs:    result.Logs,
		Cst:     "graph G { a -- b }",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func Init() {
	app := fiber.New()

	// Configure cors middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,OPTIONS",
	}))

	app.Post("/compile", HandleVisitor)

	app.Listen(":3000")
}

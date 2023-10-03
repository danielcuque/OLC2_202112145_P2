package api

import (
	E "OLC2/core/error"
	I "OLC2/core/interpreter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Response will have symbols, errors, logs and cst
type Resp struct {
	Errors    []*E.VisitorError `json:"errors"`
	Symbols   []I.ApiObject     `json:"symbols"`
	Compiled  string            `json:"compiled"`
	Optimized string            `json:"optimized"`
}

// Return errors as json

func HandleVisitor(c *fiber.Ctx) error {

	code := c.FormValue("code")

	compiler, checker := I.NewEvaluator(code)

	compiled := ""

	if !checker.HasErrors() {
		compiled = compiler.String()
	}

	response := Resp{
		Symbols:   checker.Env.GetSymbolTable(),
		Errors:    checker.GetErrors(),
		Compiled:  compiled,
		Optimized: "",
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

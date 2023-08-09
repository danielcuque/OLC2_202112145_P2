package api

import (
	"fmt"

	I "OLC2/chore/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Resp struct {
	Output  string
	Flag    bool
	Message string
}

type Message struct {
	Content string `json:"content"`
}

func HandleVisitor(c *fiber.Ctx) error {

	var message Message
	if err := c.BodyParser(&message); err != nil {
		return err
	}
	result := I.NewEvaluator(message.Content)

	out := fmt.Sprintf("%v", result.Errors)

	response := Resp{
		Output:  out,
		Flag:    true,
		Message: "Ejecución realizada con éxito",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func Init() {
	app := fiber.New()

	app.Post("/compile", HandleVisitor)

	app.Listen(":3000")
}

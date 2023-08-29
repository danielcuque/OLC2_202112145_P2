package api

import (
	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Get CST image
func GetCST(input string) string {
	apiURL := "http://lab.antlr.org/parse/"

	// parse to json lexer, grammar and input

	payload := []byte(
		fmt.Sprintf(
			`{
				"grammar": %s,
				"input": %s,
				"lexgrammar": %s,
				"start": "%s"
			}`,
			U.GetParser(),
			U.FormatInput(input),
			U.GetLexer(),
			"program",
		),
	)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}

	// create a map to store the json
	var data map[string]interface{}

	// // unmarshal the json
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	result := data["result"].(map[string]interface{})

	return result["svgtree"].(string)

}

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

	cst := GetCST(code)

	response := Resp{
		Symbols: result.Env.GetSymbolTable(),
		Errors:  result.Errors,
		Logs:    result.Logs,
		Cst:     cst,
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

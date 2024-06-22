package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CalculationRequest struct {
	N int `json:"n"`
}

type CalculationResponse struct {
	Result       uint64  `json:"result"`
	ProcessTime  float64 `json:"process_time"` // in milliseconds
}

func main() {
	app := fiber.New()

	// Serve static files (index.html and other frontend files)
	app.Static("/", "./public")

	// Endpoint to calculate Lucas number
	app.Post("/calculate", func(c *fiber.Ctx) error {
		// Parse request body
		var req CalculationRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		// Calculate Lucas number L(n)
		startTime := time.Now()
		result := Lucas(req.N)
		processTime := time.Since(startTime).Seconds() * 1000 // in milliseconds

		// Prepare response
		resp := CalculationResponse{
			Result:       result,
			ProcessTime:  processTime,
		}

		// Return JSON response
		return c.JSON(resp)
	})

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Lucas function to compute Lucas number
func Lucas(n int) uint64 {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 2
	} else if n == 1 {
		return 1
	}

	var a, b uint64 = 2, 1
	var result uint64

	for i := 2; i <= n; i++ {
		result = a + b
		a = b
		b = result
	}

	return result
}

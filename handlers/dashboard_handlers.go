package handlers

import (
	"AayushManocha/centurion/centurion-backend/middleware"
	"AayushManocha/centurion/centurion-backend/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func WeeklyDashboardHandler(c *fiber.Ctx) error {
	user, err := middleware.AuthenticatedUser(c)
	_ = user
	if err != nil {
		return c.SendString("Invalid token: " + err.Error())
	}

	date := c.Params("date")
	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid date format",
		})
	}

	// Ensure date is a Monday
	if parsedDate.Weekday() != time.Monday {
		return c.Status(400).JSON(fiber.Map{
			"error": "Date must be a Monday",
		})
	}

	categoryExpenses := services.FetchWeeklyExpensesWithCategories(user, parsedDate)

	return c.JSON(fiber.Map{
		"categoryExpenses": categoryExpenses,
	})

}

func MonthlyDashboardHandler(c *fiber.Ctx) error {
	user, err := middleware.AuthenticatedUser(c)
	_ = user
	if err != nil {
		return c.SendString("Invalid token")
	}

	date := c.Params("date")
	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid date format",
		})
	}

	// Ensure date is the first of the month
	if parsedDate.Day() != 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Date must be the first of the month",
		})
	}

	categoryExpenses := services.FetchMonthlyExpensesWithCategories(user, parsedDate)

	return c.JSON(fiber.Map{
		"categoryExpenses": categoryExpenses,
	})
}

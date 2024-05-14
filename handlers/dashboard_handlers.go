package handlers

import (
	"AayushManocha/centurion/centurion-backend/middleware"
	"AayushManocha/centurion/centurion-backend/services"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func WeeklyDashboardHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

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
	user, _ := middleware.AuthenticatedUser(c)

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

func MonthlyMetricsHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

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

	metrics := services.FetchMonthlyMetrics(user, parsedDate)

	fmt.Printf("Metrics: %+v \n", metrics)

	return c.JSON(fiber.Map{
		"metrics": metrics,
	})
}

type SignleMetricResponseDTO struct {
	Date    string
	Metrics *services.MonthlyMetric
}

type MonthlyMetricsResponseDTO []SignleMetricResponseDTO

func GetAllMonthlyMetricsHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

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

	currentDate := parsedDate
	var metrics MonthlyMetricsResponseDTO
	for i := 0; i < 12; i++ {
		metrics = append(metrics, SignleMetricResponseDTO{
			Date:    currentDate.Format("2006-01-02"),
			Metrics: services.FetchMonthlyMetrics(user, currentDate),
		})
		currentDate = currentDate.AddDate(0, -1, 0)
	}

	// metrics := services.FetchAllMonthlyMetrics(user)

	return c.JSON(fiber.Map{
		"metrics": metrics,
	})
}

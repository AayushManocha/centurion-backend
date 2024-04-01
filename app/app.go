package app

import (
	"AayushManocha/centurion/centurion-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var App *fiber.App

func InitApp() *fiber.App {
	if App != nil {
		return App
	}

	App = fiber.New()
	App.Use(cors.New())

	registerRoutes()
	return App
}

func registerRoutes() {
	// // Add access control headers
	App.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://localhost:8100")
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.Next()
	})

	App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	App.Post("onboarding/income", handlers.OnboardingIncomeHandler)
	App.Post("onboarding/spending-categories", handlers.OnboardingCategoryHandler)
	App.Get("onboarding/status", handlers.OnboardingStatusHandler)

	App.Get("dashboard/weekly/:date", handlers.WeeklyDashboardHandler)
	App.Get("dashboard/monthly/:date", handlers.MonthlyDashboardHandler)

	App.Post("expense", handlers.AddExpenseHandler)

	App.Get("categories", handlers.GetAllCategoriesHandler)
	App.Get("categories/:id", handlers.ViewCategoryHandler)

}

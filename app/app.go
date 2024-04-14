package app

import (
	"AayushManocha/centurion/centurion-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var App *fiber.App

func InitApp() *fiber.App {
	if App != nil {
		return App
	}

	App = fiber.New()
	App.Use(logger.New())
	App.Use(cors.New())

	registerRoutes()
	return App
}

func registerRoutes() {
	// // Add access control headers
	App.Use(func(c *fiber.Ctx) error {
		allowedOrigins := []string{
			"http://localhost:8100",
			"capacitor://localhost", // For iOS
			"http://localhost",      // For Android
			"192.168.2.151:8100",    //iOS emulator
		}

		_ = allowedOrigins

		origin := c.Get("Origin")
		// for _, allowedOrigin := range allowedOrigins {
		// 	if allowedOrigin == origin {
		// 		c.Set("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }
		c.Set("Access-Control-Allow-Origin", origin)
		// c.Set("Access-Control-Allow-Origin", "http://localhost:8100")
		// c.Set("Access-Control-Allow-Origin", "capacitor://localhost") // For iOS
		// c.Set("Access-Control-Allow-Origin", "http://localhost")      // For Android
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.Next()
	})

	App.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
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

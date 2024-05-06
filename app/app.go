package app

import (
	"AayushManocha/centurion/centurion-backend/handlers"
	"AayushManocha/centurion/centurion-backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var App *fiber.App

func InitApp() *fiber.App {
	// fmt.Println("Initializing app")
	if App != nil {
		return App
	}

	App = fiber.New()
	App.Use(logger.New())
	// App.Use(cors.New())
	App.Use(func(c *fiber.Ctx) error {
		// Custom PREFLIGHT request handler
		if c.Method() == "OPTIONS" {
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, X-Requested-With, Accept, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods")
			c.Set("Access-Control-Max-Age", "86400")
			return c.SendStatus(204)
		}
		return c.Next()
	})
	App.Use(func(c *fiber.Ctx) error {
		allowedOrigins := []string{
			"http://localhost:8100",
			"capacitor://localhost", // For iOS
			"http://localhost",      // For Android
			"192.168.2.151:8100",    //iOS emulator
		}

		_ = allowedOrigins

		origin := c.Get("Origin")
		c.Set("Access-Control-Allow-Origin", origin)
		// fmt.Println("Setting origin to", origin)
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header.Set("Access-Control-Allow-Origin", origin)
		return c.Next()
	})
	App.Use(middleware.EnsureAuthenticated)

	registerRoutes()
	return App
}

func registerRoutes() {
	App.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	App.Post("onboarding/income", handlers.OnboardingIncomeHandler)
	App.Post("onboarding/spending-categories", handlers.OnboardingCategoryHandler)
	App.Get("onboarding/status", handlers.OnboardingStatusHandler)

	App.Get("dashboard/weekly/:date", handlers.WeeklyDashboardHandler)
	App.Get("dashboard/monthly/:date", handlers.MonthlyDashboardHandler)
	App.Get("dashboard/metrics/monthly/:date", handlers.MonthlyMetricsHandler)

	App.Post("expense", handlers.AddExpenseHandler)

	App.Get("categories", handlers.GetAllCategoriesHandler)
	App.Get("categories/:id", handlers.ViewCategoryHandler)
	App.Delete("categories/:id", handlers.DeleteCategoryHandler)
}

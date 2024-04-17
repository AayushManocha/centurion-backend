package handlers

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

type OnboardingIncomeDTO struct {
	Income int `json:"income"`
}

func OnboardingStatusHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

	db_conn := db.GetDB()
	income_record := db.UserMonthlyIncome{}
	db_conn.Where("user_id = ?", user.ID).Find(&income_record)

	categories := []db.UserSpendingCategory{}
	db_conn.Where("user_id = ?", user.ID).Find(&categories)

	return c.JSON(fiber.Map{
		"hasIncome":           income_record.ID != 0,
		"hasSpendingCategory": len(categories) > 0,
	})

}

func OnboardingIncomeHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

	dto := new(OnboardingIncomeDTO)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	new_income_record := db.UserMonthlyIncome{
		UserID: user.ID,
		Income: dto.Income,
	}

	db_conn := db.GetDB()
	db_conn.Create(&new_income_record)

	// c.Response().Header.Set("Access-Control-Allow-Origin", "*")

	return c.JSON(fiber.Map{
		"message": "Income record created",
	})
}

type OnboardingCategories struct {
	Categories []OnboardingCategoryDTO `json:"categories"`
}

type OnboardingCategoryDTO struct {
	Title           string `json:"title"`
	BudgetAmount    int    `json:"budgetAmount"`
	IsTrackedWeekly bool   `json:"isTrackedWeekly"`
}

func OnboardingCategoryHandler(c *fiber.Ctx) error {
	user, err := middleware.AuthenticatedUser(c)

	if err != nil {
		return c.SendString("Invalid token" + err.Error())
	}

	dto := new(OnboardingCategories)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db_conn := db.GetDB()
	for _, category := range dto.Categories {
		new_category := db.UserSpendingCategory{
			UserID:          user.ID,
			Title:           category.Title,
			BudgetAmount:    category.BudgetAmount,
			IsTrackedWeekly: category.IsTrackedWeekly,
		}

		db_conn.Create(&new_category)
	}

	return c.JSON(fiber.Map{
		"message": "Saved spending categories",
	})
}

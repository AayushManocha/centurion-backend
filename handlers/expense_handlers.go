package handlers

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AddExpenseDTO struct {
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}

func AddExpenseHandler(c *fiber.Ctx) error {
	user, err := middleware.AuthenticatedUser(c)
	if err != nil {
		return c.SendString("Invalid token")
	}

	dto := new(AddExpenseDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db_conn := db.InitDB()

	// Ensure user has the category
	category := db.UserSpendingCategory{}
	db_conn.Where("user_id = ? AND id = ?", user.ID, dto.CategoryID).Find(&category)
	if category.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid category",
		})
	}

	parsedDate, _ := time.Parse("2006-01-02", dto.Date)
	expense := db.UserExpense{
		Amount:      dto.Amount,
		Date:        parsedDate,
		Description: dto.Description,
		CategoryID:  dto.CategoryID,
	}

	db_conn.Create(&expense)
	return c.JSON(fiber.Map{
		"success": true,
	})
}

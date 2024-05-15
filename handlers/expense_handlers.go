package handlers

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/middleware"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AddExpenseDTO struct {
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}

func DeleteExpenseHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

	expenseID := c.Params("id")

	fmt.Printf("Deleting expense with ID: %s\n", expenseID)
	fmt.Printf("User: %+v\n", user)

	db_conn := db.InitDB()
	var expense db.UserExpense
	db_conn.Where("id = ?", expenseID).Find(&expense)

	if expense.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid expense",
		})
	}

	db_conn.Delete(&expense)
	return c.JSON(fiber.Map{
		"success": true,
	})

}

func AddExpenseHandler(c *fiber.Ctx) error {
	user, _ := middleware.AuthenticatedUser(c)

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

	fmt.Println("adding expense to category: ", category.Title)

	parsedDate, _ := time.Parse("2006-01-02", dto.Date)
	expense := db.UserExpense{
		Amount:      dto.Amount,
		Date:        parsedDate,
		Description: dto.Description,
		CategoryID:  dto.CategoryID,
	}

	fmt.Printf("Expense: %+v \n", expense)

	db_conn.Create(&expense)
	return c.JSON(fiber.Map{
		"success": true,
	})
}

package handlers

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategoriesHandler(c *fiber.Ctx) error {
	user, err := middleware.AuthenticatedUser(c)
	if err != nil {
		return c.SendString("Invalid token")
	}

	db_conn := db.InitDB()
	var categories []db.UserSpendingCategory
	db_conn.Where("user_id = ?", user.ID).Find(&categories)

	return c.JSON(fiber.Map{
		"categories": categories,
	})
}

func ViewCategoryHandler(c *fiber.Ctx) error {
	user, err := middleware.AuthenticatedUser(c)
	if err != nil {
		return c.SendString("Invalid token")
	}

	categoryID := c.Params("id")
	db_conn := db.InitDB()
	var category db.UserSpendingCategory
	db_conn.Where("id = ? AND user_id = ?", categoryID, user.ID).First(&category)

	if category.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Category not found",
		})
	}

	var expenses []db.UserExpense
	db_conn.Where("category_id = ?", category.ID).Find(&expenses)

	return c.JSON(fiber.Map{
		"category": category,
		"expenses": expenses,
	})

}

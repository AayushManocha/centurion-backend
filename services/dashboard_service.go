package services

import (
	"AayushManocha/centurion/centurion-backend/db"
	"time"
)

type CategoryExpense struct {
	CategoryID      int    `json:"categoryId"`
	CategoryTitle   string `json:"categoryTitle"`
	TotalExpense    int    `json:"totalExpense"`
	RemainingBudget int    `json:"remainingBudget"`
	TotalBudget     int    `json:"totalBudget"`
}

func FetchWeeklyExpensesWithCategories(user db.User, date time.Time) []CategoryExpense {
	db_conn := db.GetDB()
	var weeklyCategories []db.UserSpendingCategory
	db_conn.Where("user_id = ? AND is_tracked_weekly = ?", user.ID, true).Find(&weeklyCategories)

	categoryExpenses := new([]CategoryExpense)
	for _, category := range weeklyCategories {
		var totalExpense int
		var weeklyExpenses []db.UserExpense
		db_conn.Where("category_id = ? AND date >= ? AND date <= ?", category.ID, date, date.AddDate(0, 0, 6)).Find(&weeklyExpenses)

		for _, expense := range weeklyExpenses {
			totalExpense += expense.Amount
		}

		*categoryExpenses = append(*categoryExpenses, CategoryExpense{
			CategoryID:      category.ID,
			CategoryTitle:   category.Title,
			TotalExpense:    totalExpense,
			RemainingBudget: (category.BudgetAmount / 4) - totalExpense,
			TotalBudget:     category.BudgetAmount,
		})
	}

	return *categoryExpenses
}

func FetchMonthlyExpensesWithCategories(user db.User, date time.Time) []CategoryExpense {
	db_conn := db.GetDB()
	var monthlyCategories []db.UserSpendingCategory
	db_conn.Where("user_id = ? AND is_tracked_weekly = ?", user.ID, false).Find(&monthlyCategories)

	lastDayOfMonth := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC)

	// fmt.Printf("Last day of month: %s\n", lastDayOfMonth)

	categoryExpenses := new([]CategoryExpense)
	for _, category := range monthlyCategories {
		var totalExpense int
		var monthlyExpenses []db.UserExpense
		db_conn.Where("category_id = ? AND date >= ? AND date <= ?", category.ID, date, lastDayOfMonth).Find(&monthlyExpenses)

		// fmt.Printf("Monthly expenses for category %s: %+v\n", category.Title, monthlyExpenses)

		for _, expense := range monthlyExpenses {
			totalExpense += expense.Amount
		}

		*categoryExpenses = append(*categoryExpenses, CategoryExpense{
			CategoryID:      category.ID,
			CategoryTitle:   category.Title,
			TotalExpense:    totalExpense,
			RemainingBudget: category.BudgetAmount - totalExpense,
			TotalBudget:     category.BudgetAmount,
		})
	}

	return *categoryExpenses
}

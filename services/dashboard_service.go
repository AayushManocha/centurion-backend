package services

import (
	"AayushManocha/centurion/centurion-backend/db"
	"fmt"
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

	fmt.Printf("Weekly categories: %+v \n", weeklyCategories)

	fmt.Printf("start date: %s\n", date)
	fmt.Printf("end date: %s\n", date.AddDate(0, 0, 6))

	categoryExpenses := new([]CategoryExpense)
	for _, category := range weeklyCategories {
		var totalExpense int
		var weeklyExpenses []db.UserExpense
		db_conn.Where("category_id = ? AND date >= ? AND date <= ?", category.ID, date, date.AddDate(0, 0, 6)).Find(&weeklyExpenses)

		for _, expense := range weeklyExpenses {
			totalExpense += expense.Amount
		}

		// fmt.Printf("Weekly expenses: %+v \n", weeklyExpenses)

		fmt.Printf("Total expense for category %s: %d\n", category.Title, totalExpense)

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

type MonthlyMetric struct {
	expenses    []CategoryExpense
	totalSpend  int
	totalBudget int
	remaining   int
}

func FetchMonthlyMetrics(user db.User, date time.Time) MonthlyMetric {
	db_conn := db.GetDB()
	var monthlyCategories []db.UserSpendingCategory
	db_conn.Where("user_id = ?", user.ID).Find(&monthlyCategories)

	lastDayOfMonth := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC)

	categoryExpenses := new([]CategoryExpense)
	for _, category := range monthlyCategories {
		var totalExpense int
		var monthlyExpenses []db.UserExpense
		db_conn.Where("category_id = ? AND date >= ? AND date <= ?", category.ID, date, lastDayOfMonth).Find(&monthlyExpenses)

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

	var totalSpend int
	var totalBudget int
	for _, category := range *categoryExpenses {
		totalSpend += category.TotalExpense
		totalBudget += category.TotalBudget
	}

	return MonthlyMetric{
		expenses:    *categoryExpenses,
		totalSpend:  totalSpend,
		totalBudget: totalBudget,
		remaining:   totalBudget - totalSpend,
	}
}

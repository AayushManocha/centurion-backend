package services

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/helpers"
	"testing"
	"time"
)

func TestFetchWeeklyExpensesWithCategories(t *testing.T) {
	helpers.ConfigureTests(t)

	db_conn := db.InitDB()
	var testUser db.User
	db_conn.Where("email = ?", "aayush.manocha@gmail.com").First(&testUser)

	// Create a weekly category
	weeklyCategory := db.UserSpendingCategory{
		UserID:          testUser.ID,
		Title:           "Weekly Category",
		BudgetAmount:    400,
		IsTrackedWeekly: true,
	}
	db_conn.Create(&weeklyCategory)

	//Create a monthly category
	monthlyCategory := db.UserSpendingCategory{
		UserID:          testUser.ID,
		Title:           "Monthly Category",
		BudgetAmount:    400,
		IsTrackedWeekly: false,
	}
	db_conn.Create(&monthlyCategory)

	// Create an expense for the weekly category
	expense_date, _ := time.Parse("2006-01-02", "2021-03-28")
	weeklyExpense := db.UserExpense{
		CategoryID: weeklyCategory.ID,
		Amount:     100,
		Date:       expense_date,
	}
	db_conn.Create(&weeklyExpense)

	monday, _ := time.Parse("2006-01-02", "2021-03-25")
	categoryExpenses := FetchWeeklyExpensesWithCategories(testUser, monday)

	// Print all expenses
	var expenses []db.UserExpense
	db_conn.Find(&expenses)

	if len(categoryExpenses) != 1 {
		t.Errorf("Expected 1 category, got %d", len(categoryExpenses))
	}

	if categoryExpenses[0].CategoryTitle != "Weekly Category" {
		t.Errorf("Expected category title to be 'Weekly Category', got %s", categoryExpenses[0].CategoryTitle)
	}

	if categoryExpenses[0].TotalExpense != 100 {
		t.Errorf("Expected total expense to be 100, got %d", categoryExpenses[0].TotalExpense)
	}

	if categoryExpenses[0].RemainingBudget != 0 {
		t.Errorf("Expected remaining budget to be 0, got %d", categoryExpenses[0].RemainingBudget)
	}

	if categoryExpenses[0].TotalBudget != 400 {
		t.Errorf("Expected total budget to be 400, got %d", categoryExpenses[0].TotalBudget)
	}

}

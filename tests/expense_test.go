package tests

import (
	"AayushManocha/centurion/centurion-backend/app"
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/helpers"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserCanAddExpense(t *testing.T) {
	helpers.ConfigureTests(t)

	application := app.InitApp()
	db_conn := db.InitDB()

	var user db.User
	db_conn.Where("email = ?", "aayush.manocha@gmail.com").First(&user)

	// Create test category
	test_category := db.UserSpendingCategory{
		Title:        "Test category",
		UserID:       user.ID,
		BudgetAmount: 1000,
	}
	db_conn.Create(&test_category)

	request_body := fmt.Sprintf(`{
		"amount": 100,
		"date": "2021-01-01",
		"description": "Test expense",
		"category_id": %d
	}`, test_category.ID)

	request := httptest.NewRequest("POST", "/expense", strings.NewReader(request_body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer testtoken")

	response, _ := application.Test(request)

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	// Expect expense to be created
	var expense db.UserExpense
	db_conn.Last(&db.UserExpense{}).Find(&expense)

	if expense.Amount != 100 {
		t.Errorf("Expected expense amount to be 100, but got %d", expense.Amount)
	}

	if expense.Description != "Test expense" {
		t.Errorf("Expected expense description to be 'Test expense', but got %s", expense.Description)
	}

	if expense.CategoryID != test_category.ID {
		t.Errorf("Expected expense category to be %d, but got %d", test_category.ID, expense.CategoryID)
	}

	if expense.Date.String() != "2021-01-01 00:00:00 +0000 UTC" {
		t.Errorf("Expected expense date to be '2021-01-01', but got %s", expense.Date)
	}
}

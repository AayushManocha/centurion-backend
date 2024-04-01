package tests

import (
	"AayushManocha/centurion/centurion-backend/app"
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/helpers"
	"net/http/httptest"
	"strings"
	"testing"
)

func LoginHelper() {

}

func TestOnboardingIncomeHandler(t *testing.T) {
	helpers.ConfigureTests(t)
	application := app.InitApp()

	request := httptest.NewRequest("POST", "/onboarding/income", strings.NewReader(`{"income": 1000}`))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer testtoken")
	response, _ := application.Test(request)

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	// // Expect monthly income record to be created
	db_conn := db.GetDB()

	record := db.UserMonthlyIncome{}
	db_conn.Last(&db.UserMonthlyIncome{}).Find(&record)

	if record.Income != 1000 {
		t.Errorf("Expected income to be 1000, but got %d", record.Income)
	}
}

func TestOnboardingCategoriesHandler(t *testing.T) {
	helpers.ConfigureTests(t)
	application := app.InitApp()
	request_body := `{
		"categories": [
			{"title": "Groceries", "budgetAmount": 100, "isTrackedWeekly": true},
			{"title": "Rent", "budgetAmount": 500, "isTrackedWeekly": false}
		]
	}`
	request := httptest.NewRequest("POST", "/onboarding/spending-categories", strings.NewReader(request_body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer testtoken")
	response, _ := application.Test(request)

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	// // Expect spending categories to be created

	db_conn := db.InitDB()

	var testUser db.User
	db_conn.Where("email = ?", "aayush.manocha@gmail.com").First(&testUser)

	var categories []db.UserSpendingCategory
	db_conn.Where("user_id = ?", testUser.ID).Find(&categories)

	if len(categories) != 2 {
		t.Errorf("Expected 2 spending categories to be created, but got %d", len(categories))
	}

}

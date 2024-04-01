package tests

import (
	"AayushManocha/centurion/centurion-backend/app"
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/helpers"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetAllCategories(t *testing.T) {
	helpers.ConfigureTests(t)

	application := app.InitApp()

	db_conn := db.InitDB()
	var testUser db.User
	db_conn.Where("email = ?", "aayush.manocha@gmail.com").First(&testUser)

	// Create test categories
	db_conn.Create(&db.UserSpendingCategory{
		Title:        "Test category",
		UserID:       testUser.ID,
		BudgetAmount: 1000,
	})

	db_conn.Create(&db.UserSpendingCategory{
		Title:        "Test category2",
		UserID:       testUser.ID,
		BudgetAmount: 1200,
	})

	request := httptest.NewRequest("GET", "/categories", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer testtoken")

	response, _ := application.Test(request)
	defer response.Body.Close()

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	type ResponseDTO struct {
		Categories []db.UserSpendingCategory `json:"categories"`
	}

	var responseBody ResponseDTO
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Categories) != 2 {
		t.Errorf("Expected 2 categories, but got %d", len(responseBody.Categories))
	}

	if responseBody.Categories[0].Title != "Test category" {
		t.Errorf("Expected first category to be 'Test category', but got %s", responseBody.Categories[0].Title)
	}

	if responseBody.Categories[1].Title != "Test category2" {
		t.Errorf("Expected second category to be 'Test category2', but got %s", responseBody.Categories[1].Title)
	}
}

func TestGetIndividualCategory(t *testing.T) {
	helpers.ConfigureTests(t)

	application := app.InitApp()

	db_conn := db.InitDB()
	var testUser db.User
	db_conn.Where("email = ?", "aayush.manocha@gmail.com").First(&testUser)

	// Create test category
	db_conn.Create(&db.UserSpendingCategory{
		Title:        "Test category",
		UserID:       testUser.ID,
		BudgetAmount: 1000,
	})

	// Create test expenses
	var testCategory db.UserSpendingCategory
	db_conn.Where("title = ?", "Test category").First(&testCategory)

	db_conn.Create(&db.UserExpense{
		Amount:      100,
		CategoryID:  testCategory.ID,
		Description: "Test expense",
	})

	db_conn.Create(&db.UserExpense{
		Amount:      200,
		CategoryID:  testCategory.ID,
		Description: "Test expense2",
	})

	request := httptest.NewRequest("GET", fmt.Sprintf("/categories/%d", testCategory.ID), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImNhdCI6ImNsX0I3ZDRQRDExMUFBQSIsImtpZCI6Imluc18yZUNqcU5va2sxMWF0elcxU2RhUlRBSXVLalkiLCJ0eXAiOiJKV1QifQ.eyJhenAiOiJodHRwOi8vbG9jYWxob3N0OjgxMDAiLCJleHAiOjE3MTE4NjM0NDYsImlhdCI6MTcxMTg2MzM4NiwiaXNzIjoiaHR0cHM6Ly9jbG9zZS1iYWRnZXItMTEuY2xlcmsuYWNjb3VudHMuZGV2IiwibmJmIjoxNzExODYzMzc2LCJzaWQiOiJzZXNzXzJlTGZaUFpsSDRuMmRmWkZnYkpUb2w4S05CeCIsInN1YiI6InVzZXJfMmVDbDZBUzVTZzdHdlJnWGt0ZUY4STg4Z3NuIn0.aouryvm3OD3nGRGCWgGF7Ov53MhgHMI3xrkS6gtabWZX_ip0OGtdHvtC7D-vYgLcI1THpIcjtR6cRYvAzdldcmmt-4BW-lqs3XttapCy1EiRONTU_jzOA265FCijrjsUhWEFPVFGQPrgTpL-tdEtX1afo-Rwow5xQBDEFvyRP-8Au804gL8oGosUEODjuU_SvoHMYdTVweg6kFGBOShW6g7UmJAldugQQRpK8aqu8YzoTiym78wji520zZV2Y3vUU9-ma6UsAhax-")

	response, _ := application.Test(request)

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	type ResponseDTO struct {
		Category db.UserSpendingCategory `json:"category"`
		Expenses []db.UserExpense        `json:"expenses"`
	}

	var responseBody ResponseDTO
	json.NewDecoder(response.Body).Decode(&responseBody)

	if responseBody.Category.Title != "Test category" {
		t.Errorf("Expected category title to be 'Test category', but got %s", responseBody.Category.Title)
	}

	if len(responseBody.Expenses) != 2 {
		t.Errorf("Expected 2 expenses, but got %d", len(responseBody.Category.UserExpenses))
	}

}

package helpers

import (
	"AayushManocha/centurion/centurion-backend/db"
	"testing"
)

func ConfigureTests(t *testing.T) {
	t.Setenv("ENVIRONMENT", "testing")
	t.Setenv("CLERK_SECRET_KEY", "sk_test_xOqtz2D5sJqrtVtYGeyWWHfpix1AxJBLm66acYhpFN")

	// Clear and reseed the database

	db_connection := db.InitDB()

	db_connection.Exec("DELETE FROM user_expenses")
	db_connection.Exec("DELETE FROM user_monthly_incomes")
	db_connection.Exec("DELETE FROM user_spending_categories")
	db_connection.Exec("DELETE FROM users")

	db.SeedDB()
}

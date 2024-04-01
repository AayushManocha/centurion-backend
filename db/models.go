package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int    `json:"id"`
	Email string `json:"email"`

	MonthlyIncomes []UserMonthlyIncome `json:"monthly_incomes"`
}

type UserMonthlyIncome struct {
	gorm.Model
	ID       int    `json:"id"`
	Income   int    `json:"income"`
	ActiveOn string `json:"active_on"`

	UserID int  `json:"user_id"`
	User   User `json:"user"`
}

type UserSpendingCategory struct {
	gorm.Model
	ID              int    `-`
	Title           string `json:"title"`
	BudgetAmount    int    `json:"budgetAmount"`
	IsTrackedWeekly bool   `json:"isTrackedWeekly"`

	UserID int  `json:"-"`
	User   User `json:"-"`

	UserExpenses []UserExpense `json:"expenses" gorm:"foreignKey:CategoryID"`
}

type UserExpense struct {
	gorm.Model
	ID          int       `-`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`

	CategoryID int                  `json:"category_id"`
	Category   UserSpendingCategory `json:"category"`
}

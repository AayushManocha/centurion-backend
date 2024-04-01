package db

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	fmt.Printf("Basepath: %s\n", basepath)

	var db_name string

	if os.Getenv("ENVIRONMENT") == "testing" {
		db_name = "/automated-test.db"
	} else {
		db_name = "/test.db"
	}

	db, err := gorm.Open(sqlite.Open(basepath+db_name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	SeedDB()

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserMonthlyIncome{})
	DB.AutoMigrate(&UserSpendingCategory{})
	DB.AutoMigrate(&UserExpense{})

	return DB
}

func SeedDB() {
	if os.Getenv("ENVIRONMENT") == "testing" {
		user := new(User)
		user.Email = "aayush.manocha@gmail.com"
		DB.Create(&user)
	}
}

func GetDB() *gorm.DB {
	return DB
}

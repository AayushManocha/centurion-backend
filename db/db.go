package db

import (
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/postgres"
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

	// fmt.Printf("Basepath: %s\n", basepath)

	var db_name string

	if os.Getenv("ENVIRONMENT") == "testing" {
		db_name = "/automated-test.db"
	} else {
		db_name = "/test.db"
	}

	if os.Getenv("DB_TYPE") == "postgres" {
		dsn := "postgres://aayush_manocha_centurion_db_user:6DSkE4zKvCktoPeK2VigTJQl1zAAEPtx@dpg-coqjq65jm4es73aoh4d0-a.oregon-postgres.render.com/aayush_manocha_centurion_db"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to postgres database")
		}
		DB = db
	} else {
		db, err := gorm.Open(sqlite.Open(basepath+db_name), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		DB = db
	}

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

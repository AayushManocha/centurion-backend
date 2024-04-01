package main

import (
	"AayushManocha/centurion/centurion-backend/app"
	"AayushManocha/centurion/centurion-backend/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()
	application := app.InitApp()
	application.Listen(":8080")
}

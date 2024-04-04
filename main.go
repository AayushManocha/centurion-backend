package main

import (
	"AayushManocha/centurion/centurion-backend/app"
	"AayushManocha/centurion/centurion-backend/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()
	application := app.InitApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	application.Listen(port)
}

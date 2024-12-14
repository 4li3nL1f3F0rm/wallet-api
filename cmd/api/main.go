package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/4li3nL1f3F0rm/wallet-api/internal/db"
	"github.com/4li3nL1f3F0rm/wallet-api/internal/router"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")

	dbConnection := db.Connect()
	defer dbConnection.Close()

	app := fiber.New()

	router.SetupRoutes(app, dbConnection)

	log.Printf("Starting server on port %s", PORT)
	if err := app.Listen(":" + PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

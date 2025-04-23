package main

import (
	"log"
	"myblogs/config"
	"myblogs/db"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	config.DatabaseInit()
	db.RunMigration()
	var PORT = os.Getenv("PORT")

	log.Fatal(app.Listen(":" + PORT))
}

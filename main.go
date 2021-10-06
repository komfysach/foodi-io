package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file from given path
	// we keep it empty it will load from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// getting env variables
	DB_DSN := os.Getenv("POSTGRES_DSN")

	// connecting to DB
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DB_DSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	fmt.Println("Successfully connected!")
	fmt.Println(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}

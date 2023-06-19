package main

import (
	"log"

	database "github.com/amirnajdi/order-book/Database"
	env "github.com/amirnajdi/order-book/Helper"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, err := database.Connect(database.DBconnectionConfig{
		Url:      env.Getenv("DB_URL", "localhost"),
		Port:     env.Getenv("DB_PORT", "3306"),
		Username: env.Getenv("DB_USERNAME", "root"),
		Password: env.Getenv("DB_PASSWORD", ""),
		DBname:   env.Getenv("DB_NAME", ""),
	})

	if err != nil {
		log.Fatalln(err)
	}
}

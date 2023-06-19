package main

import (
	"fmt"
	"log"

	database "github.com/amirnajdi/order-book/Database"
	env "github.com/amirnajdi/order-book/Helper"
	kafka "github.com/amirnajdi/order-book/Kafka"
	order "github.com/amirnajdi/order-book/Models"
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

	kafkaConnection := kafka.Connection()
	defer kafkaConnection.Close()

	fmt.Println("Listen for kafka data....")
	for {
		var order order.Order
		var err error
		order, err = kafka.ConsumeOrder(kafkaConnection)
		if err != nil {
			fmt.Println(err)
			continue
		}
		order.Insert()
		fmt.Println("Consome order...")
		fmt.Println(order)
	}
}

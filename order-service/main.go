package main

import (
	"log"
	"net/http"
	"os"

	"event-driven-order-system/order-service/db"
	"event-driven-order-system/order-service/handlers"
	"event-driven-order-system/order-service/kafka"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load env or not env file found")
	}

	broker := os.Getenv("BROKER")

	db.InitDB()
	kafka.InitKafkaWriter(broker, "orders")
	http.HandleFunc("/order", handlers.HandleOrder)
	http.HandleFunc("/delete", handlers.HandleDelete)
	log.Println("Order Service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

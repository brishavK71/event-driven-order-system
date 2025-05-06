package main

import (
	"log"
	"net/http"

	"event-driven-order-system/order-service/db"
	"event-driven-order-system/order-service/handlers"
)

func main() {
	db.InitDB()
	http.HandleFunc("/order", handlers.HandleOrder)
	http.HandleFunc("/delete", handlers.HandleDelete)
	log.Println("Order Service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

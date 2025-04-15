package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID       string  `json:"id"`
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

var db *gorm.DB

func initDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load env or not env file found")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&Order{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
func HandleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
	}

	if err := db.Create(&order).Error; err != nil {
		http.Error(w, "Failed to save order to DB", http.StatusInternalServerError)
	}

	log.Printf("Received order: %+v\n", order)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order Saved to the database"})
}
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
	}

	if err := db.Delete(&Order{}, order.ID).Error; err != nil {
		http.Error(w, "Failed to delete order from DB", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Order %s deleted from database", order.ID)})
}
func main() {
	initDB()
	http.HandleFunc("/order", HandleOrder)
	http.HandleFunc("/delete", HandleDelete)
	log.Println("Order Service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

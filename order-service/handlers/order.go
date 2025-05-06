package handlers

import (
	"encoding/json"
	"event-driven-order-system/order-service/db"
	"event-driven-order-system/order-service/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var bodyBytes []byte
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var orders []models.Order
	if bodyBytes[0] == '[' {

		err = json.Unmarshal(bodyBytes, &orders)
		if err != nil {
			http.Error(w, "INvalid order array data", http.StatusBadRequest)
			return
		}

		if err := db.DB.Create(&orders).Error; err != nil {
			http.Error(w, "Failed to save orders", http.StatusInternalServerError)
		}
	} else {
		var order models.Order
		err = json.Unmarshal(bodyBytes, &order)
		if err != nil {
			http.Error(w, "Invalid order data", http.StatusBadRequest)
		}

		if err := db.DB.Create(&order).Error; err != nil {
			http.Error(w, "Failed to save order", http.StatusInternalServerError)

		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Order Saved to the database"})
		}
	}

	log.Printf("Received order: %+v\n", orders)

}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
	}

	if err := db.DB.Delete(&models.Order{}, order.ID).Error; err != nil {
		http.Error(w, "Failed to delete order from DB", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Order %s deleted from database", order.ID)})
	}
}

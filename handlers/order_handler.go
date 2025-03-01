package handlers

import (
	"database/sql"
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"ecommerce-backend/queue"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if order.OrderID == 0 {
		http.Error(w, "order_id is required", http.StatusBadRequest)
		return
	}

	order.Status = "Pending"
	order.CreatedTime = time.Now().Format("2006-01-02 15:04:05")

	queue.OrderQueue <- order

	response := map[string]string{"message": "Order processing initiated successfully", "order_id": strconv.Itoa(order.OrderID)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func GetOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["order_id"])

	var order models.Order
	var itemIDsStr, createdTime, processedTime sql.NullString

	err := config.DB.QueryRow("SELECT order_id, user_id, item_ids, total_amount, status, created_time, processed_time FROM orders WHERE order_id = ?",
		orderID).Scan(&order.OrderID, &order.UserID, &itemIDsStr, &order.TotalAmount, &order.Status, &createdTime, &processedTime)

	if err == sql.ErrNoRows {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(itemIDsStr.String), &order.ItemIDs)
	if err != nil {
		http.Error(w, "Error parsing item IDs", http.StatusInternalServerError)
		return
	}

	order.CreatedTime = createdTime.String
	order.ProcessedTime = processedTime.String

	response, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

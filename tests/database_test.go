package tests

import (
	"ecommerce-backend/config"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	config.InitDB()
	if config.DB == nil {
		t.Fatal("Database connection failed")
	}
}

func TestInsertOrder(t *testing.T) {
	config.InitDB()
	_, err := config.DB.Exec("INSERT INTO orders (order_id, user_id, item_ids, total_amount, status, created_time, processed_time) VALUES (?, ?, ?, ?, ?, ?, ?)",
		1234567, 1, "[101, 102]", 50.0, "Pending", "2025-03-01 12:00:00", "")
	if err != nil {
		t.Fatalf("Failed to insert order: %v", err)
	}
}

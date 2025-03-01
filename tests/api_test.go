package tests

import (
	"bytes"
	"ecommerce-backend/config"
	"ecommerce-backend/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateOrderHandler(t *testing.T) {
	config.InitDB()
	reqBody := []byte(`{"order_id": 1, "user_id": 1, "item_ids": [101, 102], "total_amount": 50.0}`)
	req, err := http.NewRequest("POST", "/order", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}
}

func TestGetOrderStatusHandler(t *testing.T) {
	config.InitDB()
	req, err := http.NewRequest("GET", "/order/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/order/{order_id}", handlers.GetOrderStatusHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

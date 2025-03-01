package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/handlers"
	"ecommerce-backend/queue"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	go queue.ProcessOrders()

	r := mux.NewRouter()
	r.HandleFunc("/order", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/order/{order_id}", handlers.GetOrderStatusHandler).Methods("GET")
	r.HandleFunc("/metrics", handlers.GetMetricsHandler).Methods("GET")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  300 * time.Second,
	}

	log.Println("Server started on :8080")
	log.Fatal(server.Serve(listener))
}

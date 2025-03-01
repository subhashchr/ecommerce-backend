package handlers

import (
	"database/sql"
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"encoding/json"
	"net/http"
)

func GetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	var totalOrders, pending, processing, completed int
	var totalProcessingTime float64
	var completedOrdersCount int

	rows, err := config.DB.Query("SELECT status, created_time, processed_time FROM orders")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		var createdTime, processedTime sql.NullTime

		err := rows.Scan(&status, &createdTime, &processedTime)
		if err != nil {
			continue
		}

		totalOrders++

		switch status {
		case "Pending":
			pending++
		case "Processing":
			processing++
		case "Completed":
			completed++
			if createdTime.Valid && processedTime.Valid {
				totalProcessingTime += processedTime.Time.Sub(createdTime.Time).Seconds()
				completedOrdersCount++
			}
		}
	}

	avgProcessingTime := 0.0
	if completedOrdersCount > 0 {
		avgProcessingTime = totalProcessingTime / float64(completedOrdersCount)
	}

	metrics := models.Metrics{
		TotalOrders:      totalOrders,
		AvgProcessing:    avgProcessingTime,
		PendingOrders:    pending,
		ProcessingOrders: processing,
		CompletedOrders:  completed,
	}

	response, _ := json.Marshal(metrics)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	baseURL   = "http://localhost:8080/order"
	numOrders = 1000
)

var failedRequests int32 // Atomic counter for failed requests

func sendOrder(orderID int, client *http.Client) {
	orderJSON := fmt.Sprintf(`{"order_id": %d, "user_id": 1, "item_ids": [101, 102], "total_amount": 50.0}`, orderID)
	req, _ := http.NewRequest("POST", baseURL, bytes.NewBufferString(orderJSON))
	req.Header.Set("Content-Type", "application/json")

	for retries := 1; retries <= 10; retries++ {
		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			return // Success
		}

		time.Sleep(time.Duration(500*(retries+1)) * time.Millisecond) // Exponential backoff

	}
	fmt.Printf("Failed to create order %d after retries\n", orderID)
	atomic.AddInt32(&failedRequests, 1) // Increment failed request counter
}

func main() {
	var wg sync.WaitGroup
	client := &http.Client{Timeout: 10 * time.Second}

	for i := 1; i <= numOrders; i += 1 {
		wg.Add(1)
		go func(orderID int) {
			defer wg.Done()
			sendOrder(orderID, client)
		}(i)
	}

	wg.Wait()

	// Log failed requests
	fmt.Printf("Load test completed: %d total requests, %d failed requests\n", numOrders, atomic.LoadInt32(&failedRequests))
}

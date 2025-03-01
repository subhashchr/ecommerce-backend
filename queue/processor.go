package queue

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

var OrderQueue = make(chan models.Order, 1000)
var wg sync.WaitGroup

func ProcessOrders() {
	workers := 10

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for order := range OrderQueue {
				order.Status = "Completed"
				order.ProcessedTime = time.Now().Format("2006-01-02 15:04:05")

				for retries := 0; retries < 10; retries++ {
					_, err := config.DB.Exec("INSERT INTO orders (order_id, user_id, item_ids, total_amount, status, created_time, processed_time) VALUES (?, ?, ?, ?, ?, ?, ?)",
						order.OrderID, order.UserID, strings.Join(strings.Fields(fmt.Sprint(order.ItemIDs)), ","), order.TotalAmount, order.Status, order.CreatedTime, order.ProcessedTime)
					if err == nil {
						fmt.Printf("Order %d processed at %s, created at %s", order.OrderID, order.ProcessedTime, order.CreatedTime)
						break
					}

					if strings.Contains(err.Error(), "database is locked") {
						fmt.Printf("Database locked, retrying order %d in %d ms...", order.OrderID, 200*(retries+1))
						time.Sleep(time.Duration(200*(retries+1)) * time.Millisecond)
					} else {
						log.Println("Error inserting order:", err)
						break
					}
				}
			}
		}()
	}
}

func WaitForCompletion() {
	close(OrderQueue)
	wg.Wait()
}

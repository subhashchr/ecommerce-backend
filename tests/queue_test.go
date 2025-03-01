package tests

import (
    "ecommerce-backend/models"
    "ecommerce-backend/queue"
    "testing"
)

func TestQueueProcessing(t *testing.T) {
    order := models.Order{OrderID: 2001, UserID: 2, ItemIDs: []int{201, 202}, TotalAmount: 75.0, Status: "Pending"}
    queue.OrderQueue <- order
    close(queue.OrderQueue)

    go queue.ProcessOrders()

    if order.Status != "Completed" {
        t.Errorf("Order not processed correctly, got status %s", order.Status)
    }
}

package models

type Order struct {
    OrderID       int    `json:"order_id"`
    UserID        int    `json:"user_id"`
    ItemIDs       []int  `json:"item_ids"`
    TotalAmount   float64 `json:"total_amount"`
    Status        string `json:"status"`
    CreatedTime   string `json:"created_time"`
    ProcessedTime string `json:"processed_time,omitempty"`
}

package models

type Metrics struct {
    TotalOrders    int     `json:"total_orders"`
    AvgProcessing  float64 `json:"avg_processing_time"`
    PendingOrders  int     `json:"pending_orders"`
    ProcessingOrders int   `json:"processing_orders"`
    CompletedOrders int    `json:"completed_orders"`
}

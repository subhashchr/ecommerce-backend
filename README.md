# E-Commerce Backend System

## 📌 Overview
This project implements a **scalable backend system** for managing orders in an **e-commerce platform**. The system:
- Provides **RESTful APIs** for order management.
- Uses an **asynchronous queue** for order processing.
- Supports **SQLite** as the database with optimized write handling.
- Implements **batch processing** to handle **1,000 concurrent orders** efficiently.

## 📌 Setup Instructions

### 1️⃣ Prerequisites
Ensure you have:
- **Go 1.18+** installed → [Download Go](https://golang.org/dl/)
- **SQLite** installed → [Download SQLite](https://www.sqlite.org/download.html)

### 2️⃣ Clone the Repository
```sh
git clone <GITHUB_REPO_LINK>
cd ecommerce-backend
```

### 3️⃣ Install Dependencies
```sh
go mod tidy
```

### 4️⃣ Initialize the Database
```sh
sqlite3 orders.db < database/schema.sql
```

### 5️⃣ Run the Application
```sh
go run main.go
```

## 📌 Running the Load Test
To simulate **1,000 concurrent orders**:
```sh
go run load_test/generate_load.go
```

## 📌 API Endpoints

| Method | Endpoint | Description |
|--------|---------|-------------|
| **POST** | `/order` | Create an order (async processing) |
| **GET** | `/order/{order_id}` | Get the status of an order |
| **GET** | `/metrics` | Fetch system metrics |

## 📌 Example API Requests

### 1️⃣ Create Order
```sh
curl -X POST http://localhost:8080/order      -H "Content-Type: application/json"      -d '{
           "order_id": 12345,
           "user_id": 1,
           "item_ids": [101, 102, 103],
           "total_amount": 99.99
         }'
```
✅ **Response:**
```json
{
  "message": "Order processing initiated successfully",
  "order_id": "12345"
}
```

### 2️⃣ Get Order Status
```sh
curl -X GET http://localhost:8080/order/12345
```
✅ **Response:**
```json
{
  "order_id": 12345,
  "user_id": 1,
  "item_ids": [101, 102, 103],
  "total_amount": 99.99,
  "status": "Completed",
  "created_time": "2025-03-01 12:00:00",
  "processed_time": "2025-03-01 12:00:10"
}
```

### 3️⃣ Get System Metrics
```sh
curl -X GET http://localhost:8080/metrics
```
✅ **Response:**
```json
{
  "total_orders": 1000,
  "avg_processing_time": 5.2,
  "pending_orders": 0,
  "processing_orders": 0,
  "completed_orders": 1000
}
```

## 📌 Design Decisions & Trade-offs

### ✅ Why SQLite?
- Chosen for **simplicity and lightweight nature**.
- Supports **WAL mode for better concurrency**.
- **Limitation:** Only one write operation at a time.

### ✅ Asynchronous Processing
- Orders are **pushed to a queue** and processed **in background workers**.
- **Trade-off:** Immediate confirmation to users, but they need to check order status separately.

### ✅ Batch Load Testing
- Load test script sends requests in **controlled batches**.
- **Trade-off:** Helps prevent server overload but might delay some orders.

## 📌 Assumptions
1. **Orders are uniquely identified by `order_id`**, which is provided by the client.
2. **Item details are stored as an array of item IDs** rather than a separate table.
3. **The system processes orders as fast as possible**, but **database locks are retried** in case of `SQLITE_BUSY` errors.
4. **Processing time is calculated based on `created_time` and `processed_time`**.

## 📌 GitHub Repository
The complete code is available at:  
📌 **[GitHub Repo: <YOUR_GITHUB_LINK>]**

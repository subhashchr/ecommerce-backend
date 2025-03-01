-- Insert sample data into orders table

INSERT INTO orders (order_id, user_id, item_ids, total_amount, status, created_time, processed_time) VALUES
(1, 101, '[201, 202]', 49.99, 'Completed', '2025-03-01 10:00:00', '2025-03-01 10:05:00'),
(2, 102, '[203, 204]', 79.50, 'Pending', '2025-03-01 10:10:00', NULL),
(3, 103, '[205, 206]', 120.00, 'Processing', '2025-03-01 10:15:00', NULL),
(4, 104, '[207, 208, 209]', 89.75, 'Completed', '2025-03-01 10:20:00', '2025-03-01 10:30:00'),
(5, 105, '[210, 211]', 55.25, 'Pending', '2025-03-01 10:35:00', NULL);

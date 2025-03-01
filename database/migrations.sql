CREATE TABLE IF NOT EXISTS orders (
        order_id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        item_ids TEXT,
        total_amount REAL,
        status TEXT,
        created_time DATETIME,
        processed_time DATETIME
    );

package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "orders.db?_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`PRAGMA journal_mode = WAL;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS orders (
        order_id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        item_ids TEXT,
        total_amount REAL,
        status TEXT,
        created_time TEXT,
        processed_time TEXT
    )`)
	if err != nil {
		log.Fatal(err)
	}
}

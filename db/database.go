package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
    db, err := sql.Open("sqlite3", "charges.db")
    if err != nil {
        log.Fatalf("Failed to connect to SQLite database: %v", err)
    }

    // Ensure the charges table exists
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS charges (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        concept TEXT NOT NULL,
        amount REAL NOT NULL,
        payment_method TEXT NOT NULL,
        category TEXT NOT NULL,
        date TEXT NOT NULL
    );
    `
    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatalf("Failed to create charges table: %v", err)
    }

    return db
}

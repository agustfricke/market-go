package database

import (
	"database/sql"
	"fmt"
	"log"
)

func createUserTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
      id VARCHAR(36) PRIMARY KEY,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
      username VARCHAR(100), 
			password VARCHAR(100) NOT NULL,
			verified BOOLEAN NOT NULL DEFAULT false,
			social_id VARCHAR(100),
			avatar VARCHAR(255) NOT NULL DEFAULT 'default.png',
			otp_enabled BOOLEAN DEFAULT false,
			otp_verified BOOLEAN DEFAULT false,
			otp_secret VARCHAR(255),
			otp_auth_url VARCHAR(255)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println("user table created")
}

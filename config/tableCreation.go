package config

import (
	"database/sql"
	"fmt"
	"log"
)


func CreateUserTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			username TEXT,
			password TEXT
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User table created successfully")

}
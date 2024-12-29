package config

import (
	"database/sql"
	"log"
)

func GetConnexionClient() *sql.DB{
	connStr :="postgres://postgres:postgres@localhost:5431/gopgtest?sslmode=disable"
	db,err:=sql.Open("postgres",connStr)

	// connect to the database
	if(err!=nil){
		
		log.Fatal(err)
	}
	if(db.Ping()!=nil){
		log.Fatal(err)
	}
	return db
}
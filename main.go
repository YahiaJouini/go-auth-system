package main

import (
	"fmt"
	"log"
	"net/http"
	"server/routes"

	// _ will import the package and call the init function
	_ "github.com/lib/pq"

	"server/config"
)


func main(){

	
	db:=config.GetConnexionClient()


	// The deferred function call is executed at the end of the function where it is declared.
	// meaning db.Close() will be called at the end of the main function
	defer db.Close()

	// Creating the tables
	config.CreateUserTable(db)

	// Initialize the routes
	router:=routes.InitializeRoutes()


	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}
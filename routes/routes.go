package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router{
	router:=mux.NewRouter()
	router.HandleFunc("/auth/login",controllers.Login).Methods("POST")
	router.HandleFunc("/auth/register",controllers.Register).Methods("POST")

	// protected routes
	protected := router.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/users",controllers.GetUsers).Methods("GET")
	
	return router
}
package middleware

import (
	"fmt"
	"net/http"
	"server/config"
	"server/helpers"
	"strings"
)


func AuthMiddleware(next http.Handler)http.Handler{
	fmt.Println("Middleware running")
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			helpers.RespondWithError(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}
	
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			helpers.RespondWithError(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}
	
		err :=config.VerifyToken(tokenParts[1])
		if err != nil {
			helpers.RespondWithError(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	
	})
}


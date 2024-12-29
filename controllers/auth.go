package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/config"
	"strings"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}


var db =config.GetConnexionClient()

func Register(w http.ResponseWriter, r *http.Request) {
	// type of response 
	w.Header().Set("Content-Type", "application/json")

	// Decode the request payload into a go understandable struct
	var payload User
	decodeErr := json.NewDecoder(r.Body).Decode(&payload)
	if decodeErr != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var hashedPassowrd,hashErr = bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if(hashErr!=nil){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "An error occured",
		})
	}

	query := `
		INSERT INTO users(username, password)
		VALUES($1, $2)
`
	var _, queryErr = db.Exec(query, payload.Username, hashedPassowrd)
	if queryErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User registration failed"))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user": map[string]string{
			"username": payload.Username,
		},
	})
}

func Login(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var payload User
	decodeErr := json.NewDecoder(r.Body).Decode(&payload)
	if decodeErr != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if payload.Username == "" || payload.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Username or password is missing",
		})
		return
	}

	payload.Username = strings.TrimSpace(payload.Username)
    payload.Username = strings.ToLower(payload.Username)
	

	query := `SELECT password FROM users WHERE username=$1`
	var hashedPassword string
    err := db.QueryRow(query, payload.Username).Scan(&hashedPassword)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid credentials",
		})
		return
	}
	

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password))

	if(err!=nil){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Credentials do not match",
		})
		return
	}

	token,err:=config.GenerateToken(payload.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Error generating token",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged in successfully",
		"token":   token,
	})
}
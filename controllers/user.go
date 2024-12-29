package controllers

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/helpers"
)


func GetUsers(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	db:=config.GetConnexionClient()
	query:=`SELECT * FROM users`

	rows,err:=db.Query(query)
	if err!=nil{
		helpers.RespondWithError(w,"Error fetching users",http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next(){
		var user User
		err:=rows.Scan(&user.ID,&user.Username,&user.Password)
		if err!=nil{
			helpers.RespondWithError(w,"Error fetching users",http.StatusInternalServerError)
			return
		}
		users=append(users,user)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
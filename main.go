package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var users []User

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)

	json.NewEncoder(w).Encode(user)

	if user.Email != "c137@onecause.com" || user.Password != "#th@nH@rm#y#r!$100%D0p#" {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	} else {
		return
	}
}

func main() {

	r := mux.NewRouter()
	users = append(users, User{})
	handler := cors.Default().Handler(r)
	r.HandleFunc("/login", login).Methods("POST", "OPTONS")
	log.Fatal(http.ListenAndServe(":8080", handler))

}

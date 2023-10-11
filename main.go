package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/profile", ProfileHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Register user in DB
	// On success return JWT
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verify user credentials from DB
	// On success return JWT
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Verify JWT token
	// Return user profile
}

package main

import (
	"encoding/json"
	"net/http"

	"bookmysalon/pkg/database"
	"bookmysalon/pkg/jwt"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/profile", ProfileHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // This should ideally be hashed, not plain text.
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;"
	if err := db.QueryRow(query, u.Username, hashedPassword).Scan(&u.ID); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	token, err := jwt.GenerateToken(u.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var dbUser User
	query := "SELECT id, password FROM users WHERE username=$1;"
	if err := db.QueryRow(query, u.Username).Scan(&dbUser.ID, &dbUser.Password); err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateToken(u.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	claims, err := jwt.VerifyToken(tokenHeader)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u User
	query := "SELECT id, username FROM users WHERE username=$1;"
	if err := db.QueryRow(query, claims.Username).Scan(&u.ID, &u.Username); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	u.Password = "" // Do not send the password, even if hashed
	if response, err := json.Marshal(u); err != nil {
		http.Error(w, "Failed to process user data", http.StatusInternalServerError)
		return
	} else {
		w.Write(response)
	}
}

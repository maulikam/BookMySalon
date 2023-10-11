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

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // This should ideally be hashed, not plain text.
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Hash the password (for security)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Insert into DB
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;"
	err = db.QueryRow(query, u.Username, hashedPassword).Scan(&u.ID)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := GenerateToken(u.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Verify credentials from DB
	var dbUser User
	query := "SELECT id, password FROM users WHERE username=$1;"
	err = db.QueryRow(query, u.Username).Scan(&dbUser.ID, &dbUser.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := GenerateToken(u.Username)
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

	claims, err := VerifyToken(tokenHeader)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	db := connect()
	defer db.Close()

	// Fetch user details
	var u User
	query := "SELECT id, username FROM users WHERE username=$1;"
	err = db.QueryRow(query, claims.Username).Scan(&u.ID, &u.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Do not send the password, even if hashed
	u.Password = ""
	response, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Failed to process user


package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"bookmysalon/pkg/jwt"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

const (
	ServerAddress           = ":8080"
	BcryptCostFactor        = 8
	DatabaseErrorMessage    = "Database error"
	BadRequestMessage       = "Bad request"
	TokenErrorMessage       = "Failed to generate token"
	HashingErrorMessage     = "Failed to hash password"
	UserNotFoundMessage     = "User not found"
	InvalidPasswordMessage  = "Invalid password"
	InvalidTokenMessage     = "Invalid token"
	MissingTokenMessage     = "Missing token"
	ProcessingDataErrorMess = "Failed to process user data"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/profile", ProfileHandler).Methods("GET")

	http.ListenAndServe(ServerAddress, r)
}

func connectToDatabase() (*sql.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDatabase()
	if err != nil {
		http.Error(w, DatabaseErrorMessage, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := registerUser(db, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := jwt.GenerateToken(u.Username)
	if err != nil {
		http.Error(w, TokenErrorMessage, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func registerUser(db *sql.DB, u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), BcryptCostFactor)
	if err != nil {
		return errors.New(HashingErrorMessage)
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;"
	if err := db.QueryRow(query, u.Username, hashedPassword).Scan(&u.ID); err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDatabase()
	if err != nil {
		http.Error(w, DatabaseErrorMessage, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	token, err := loginUser(db, &u)
	if err != nil {
		switch err.Error() {
		case UserNotFoundMessage, InvalidPasswordMessage:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(token))
}

func loginUser(db *sql.DB, u *models.User) (string, error) {
	var dbUser models.User
	query := "SELECT id, password FROM users WHERE username=$1;"
	if err := db.QueryRow(query, u.Username).Scan(&dbUser.ID, &dbUser.Password); err != nil {
		return "", errors.New(UserNotFoundMessage)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)); err != nil {
		return "", errors.New(InvalidPasswordMessage)
	}

	token, err := jwt.GenerateToken(u.Username)
	if err != nil {
		return "", errors.New(TokenErrorMessage)
	}
	return token, nil
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		http.Error(w, MissingTokenMessage, http.StatusUnauthorized)
		return
	}

	claims, err := jwt.VerifyToken(tokenHeader)
	if err != nil {
		http.Error(w, InvalidTokenMessage, http.StatusUnauthorized)
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		http.Error(w, DatabaseErrorMessage, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	userProfile, err := fetchUserProfile(db, claims.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(userProfile)
	if err != nil {
		http.Error(w, ProcessingDataErrorMess, http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func fetchUserProfile(db *sql.DB, username string) (*models.User, error) {
	var u models.User
	query := "SELECT id, username FROM users WHERE username=$1;"
	if err := db.QueryRow(query, username).Scan(&u.ID, &u.Username); err != nil {
		return nil, errors.New(UserNotFoundMessage)
	}
	u.Password = "" // Ensure password is not exposed
	return &u, nil
}

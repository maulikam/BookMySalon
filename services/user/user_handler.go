// Package classification of Book My Salon API.
//
// Documentation for Book My Salon API.
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"bookmysalon/pkg/jwt"
)

const (
	ServerAddress           = ":8080"
	BcryptCostFactor        = 8
	DatabaseErrorMessage    = "Database error"
	BadRequestMessage       = "Bad request"
	TokenErrorMessage       = "failed to generate token"
	HashingErrorMessage     = "failed to hash password"
	UserNotFoundMessage     = "user not found"
	InvalidPasswordMessage  = "invalid password"
	InvalidTokenMessage     = "Invalid token"
	MissingTokenMessage     = "Missing token"
	ProcessingDataErrorMess = "Failed to process user data"
)

type UserHandler struct {
	UserService UserService // Assuming UserService is the interface that UserServiceImpl implements.
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

// Register a new user
// swagger:route POST /register users registerUser
//
// Responses:
//
//	200: tokenResponse
//	400: errorResponse
//	500: errorResponse
func (handler *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
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

	if err := handler.UserService.RegisterUser(db, &u); err != nil {
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

// User login
// swagger:route POST /login users loginUser
//
// Responses:
//
//	200: tokenResponse
//	401: errorResponse
//	500: errorResponse
func (handler *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch database connection
	db, err := database.Connect()
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

	token, err := handler.UserService.LoginUser(db, &u)
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

// Assuming there's a helper function to fetch claims from context:
func getUserClaimsFromContext(r *http.Request) (jwt.Claims, error) {
	claims, ok := r.Context().Value("userClaims").(jwt.Claims)
	if !ok {
		return claims, errors.New("no user claims found in context")
	}
	return claims, nil
}

// Get user profile
// swagger:route GET /profile users userProfile
//
// Responses:
//
//	200: userResponse
//	401: errorResponse
//	404: errorResponse
//	500: errorResponse
func (handler *UserHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := getUserClaimsFromContext(r)
	if err != nil {
		http.Error(w, InvalidTokenMessage, http.StatusUnauthorized)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, DatabaseErrorMessage, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	userProfile, err := handler.UserService.FetchUserProfile(db, claims.Username)
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

// swagger:route PUT /profile users updateUserProfile
//
// Responses:
//
//	200: userResponse
//	401: errorResponse
//	400: errorResponse
//	500: errorResponse
func (handler *UserHandler) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := getUserClaimsFromContext(r)
	if err != nil {
		http.Error(w, InvalidTokenMessage, http.StatusUnauthorized)
		return
	}

	db, err := database.Connect()
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

	u.Username = claims.Username // Ensure the username from token is used
	if err := handler.UserService.UpdateUserProfile(db, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Profile updated successfully"))
}

// swagger:route PUT /change-password users changePassword
//
// Responses:
//
//	200: messageResponse
//	401: errorResponse
//	400: errorResponse
//	500: errorResponse
func (handler *UserHandler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := getUserClaimsFromContext(r)
	if err != nil {
		http.Error(w, InvalidTokenMessage, http.StatusUnauthorized)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, DatabaseErrorMessage, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var pcr PasswordChangeRequest
	if err := json.NewDecoder(r.Body).Decode(&pcr); err != nil {
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := handler.UserService.ChangeUserPassword(db, claims.Username, pcr.OldPassword, pcr.NewPassword); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Password changed successfully"))
}

// swagger:route DELETE /profile users deleteUser
//
// Responses:
//
//	200: messageResponse
//	401: errorResponse
//	500: errorResponse
func (handler *UserHandler) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := getUserClaimsFromContext(r)
	if err != nil {
		http.Error(w, InvalidTokenMessage, http.StatusUnauthorized)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, DatabaseErrorMessage, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := handler.UserService.DeleteUserAccount(db, claims.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Account deleted successfully"))
}

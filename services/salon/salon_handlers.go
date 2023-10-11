package salon

// go-swagger annotations
// @title Salon API
// @version 1.0
// @description This is the salon service API documentation.
// @host localhost:8080
// @BasePath /
// @schemes http

import (
	"bookmysalon/models"
	"bookmysalon/pkg/jwt"
	"encoding/json"
	"net/http"
)

type SalonHandler struct {
	service SalonService
}

func NewSalonHandler(s SalonService) *SalonHandler {
	return &SalonHandler{service: s}
}

// @Summary Create a new salon
// @Description Create a new salon with the input payload
// @Accept  json
// @Produce  json
// @Param salon body models.Salon true "Create salon"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /salon [post]
func (h *SalonHandler) CreateSalon(w http.ResponseWriter, r *http.Request) {
	var salon models.Salon

	if err := json.NewDecoder(r.Body).Decode(&salon); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	salonID, err := h.service.AddSalon(salon)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"salon_id": salonID})
}

// @Summary Update salon details
// @Description Update details of an existing salon
// @Accept  json
// @Produce  json
// @Param salon body models.Salon true "Update Salon"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /salon/update [put]
func (h *SalonHandler) UpdateSalonDetails(w http.ResponseWriter, r *http.Request) {
	var salon models.Salon

	if err := json.NewDecoder(r.Body).Decode(&salon); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateSalon(salon); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := authHeader[len("Bearer "):]

		_, err := jwt.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

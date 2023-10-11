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
	"strconv"

	"github.com/gorilla/mux"
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

// @Summary Delete a salon
// @Description Delete a salon by ID
// @Accept  json
// @Produce  json
// @Param salonID path int true "Salon ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string "Salon Not Found"
// @Failure 500 {object} map[string]string
// @Router /salon/{salonID} [delete]
func (h *SalonHandler) DeleteSalon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteSalon(salonID); err != nil {
		// Depending on the nature of the error, you might want to differentiate
		// between a "Salon Not Found" error and other types of errors.
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Get salon details
// @Description Get details of a salon by ID
// @Accept  json
// @Produce  json
// @Param salonID path int true "Salon ID"
// @Success 200 {object} models.Salon
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string "Salon Not Found"
// @Failure 500 {object} map[string]string
// @Router /salon/{salonID} [get]
func (h *SalonHandler) GetSalonDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	salon, err := h.service.GetSalonByID(salonID)
	if err != nil {
		switch err {
		case ErrSalonNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Send the retrieved salon details as JSON
	json.NewEncoder(w).Encode(salon)
}

// @Summary List all salons
// @Description Retrieve a list of all salons
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Salon
// @Failure 500 {object} map[string]string
// @Router /salons [get]
func (h *SalonHandler) ListAllSalons(w http.ResponseWriter, r *http.Request) {
	salons, err := h.service.ListSalons()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send the retrieved salons as JSON
	json.NewEncoder(w).Encode(salons)
}

// @Summary Add a new service
// @Description Create a new service with the input payload
// @Accept  json
// @Produce  json
// @Param service body models.Service true "Create Service"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service [post]
func (h *SalonHandler) AddService(w http.ResponseWriter, r *http.Request) {
	var service models.Service

	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	serviceID, err := h.service.AddService(service)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"service_id": serviceID})
}

// @Summary Update service details
// @Description Update details of an existing service
// @Accept  json
// @Produce  json
// @Param service body models.Service true "Update Service"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service/update [put]
func (h *SalonHandler) UpdateServiceDetails(w http.ResponseWriter, r *http.Request) {
	var service models.Service

	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateService(service); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Delete a service
// @Description Delete a service by its ID
// @Accept  json
// @Produce  json
// @Param serviceID path int true "Service ID"
// @Success 200
// @Failure 400 {object} map[string]string "Invalid Service ID"
// @Failure 500 {object} map[string]string
// @Router /service/{serviceID} [delete]
func (h *SalonHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID, err := strconv.Atoi(vars["serviceID"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteService(serviceID); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Get a service by ID
// @Description Retrieve details of a service by its ID
// @Accept  json
// @Produce  json
// @Param serviceID path int true "Service ID"
// @Success 200 {object} models.Service
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service/{serviceID} [get]
func (h *SalonHandler) GetServiceDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID, err := strconv.Atoi(vars["serviceID"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	service, err := h.service.GetServiceByID(serviceID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(service)
}

// @Summary List all services by a salon
// @Description Retrieve all services offered by a specific salon
// @Accept  json
// @Produce  json
// @Param salonID path int true "Salon ID"
// @Success 200 {array} models.Service
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /salon/{salonID}/services [get]
func (h *SalonHandler) GetServicesBySalon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	services, err := h.service.ListServicesBySalon(salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(services)
}

// @Summary Get the average rating of a salon
// @Description Retrieve the average rating of a salon by its ID
// @Accept  json
// @Produce  json
// @Param salonID path int true "Salon ID"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /salon/{salonID}/average-rating [get]
func (h *SalonHandler) GetSalonAverageRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	avgRating, err := h.service.GetAverageRating(salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"average_rating": avgRating})
}

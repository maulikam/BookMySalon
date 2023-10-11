package availability

import (
	"bookmysalon/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type AvailabilityHandler struct {
	service AvailabilityService
}

func NewAvailabilityHandler(s AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{service: s}
}

// @Summary Create a new availability
// @Description Create a new availability with the input payload
// @Accept  json
// @Produce  json
// @Param availability body models.Availability true "Create Availability"
// @Success 201 {object} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availability [post]
func (h *AvailabilityHandler) CreateAvailability(w http.ResponseWriter, r *http.Request) {
	var availability models.Availability

	if err := json.NewDecoder(r.Body).Decode(&availability); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newAvailability, err := h.service.CreateAvailability(&availability)
	if err != nil {
		log.Println("Failed to create availability:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAvailability)
}

// @Summary Get availability details
// @Description Get details of an availability by ID
// @Accept  json
// @Produce  json
// @Param availabilityID path int true "Availability ID"
// @Success 200 {object} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string "Availability Not Found"
// @Failure 500 {object} map[string]string
// @Router /availability/{availabilityID} [get]
func (h *AvailabilityHandler) GetAvailabilityDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	availabilityID, err := strconv.Atoi(vars["availabilityID"])
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	availability, err := h.service.GetAvailabilityByID(availabilityID)
	if err != nil {
		switch err {
		case ErrAvailabilityNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(availability)
}

// @Summary Update availability details
// @Description Update details of an existing availability
// @Accept  json
// @Produce  json
// @Param availability body models.Availability true "Update Availability"
// @Success 200 {object} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availability/update [put]
func (h *AvailabilityHandler) UpdateAvailabilityDetails(w http.ResponseWriter, r *http.Request) {
	var availability models.Availability

	if err := json.NewDecoder(r.Body).Decode(&availability); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	updatedAvailability, err := h.service.UpdateAvailability(&availability)
	if err != nil {
		log.Println("Failed to update availability:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedAvailability)
}

// @Summary Delete an availability
// @Description Delete an availability by its ID
// @Accept  json
// @Produce  json
// @Param availabilityID path int true "Availability ID"
// @Success 204 "Successfully deleted"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string "Availability Not Found"
// @Failure 500 {object} map[string]string
// @Router /availability/{availabilityID} [delete]
func (h *AvailabilityHandler) DeleteAvailability(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	availabilityID, err := strconv.Atoi(vars["availabilityID"])
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteAvailability(availabilityID)
	if err != nil {
		switch err {
		case ErrAvailabilityNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary List availabilities by salon ID
// @Description Retrieve all availabilities of a specific salon by its salon ID
// @Accept  json
// @Produce  json
// @Param salonID path int true "Salon ID"
// @Success 200 {array} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availabilities/salon/{salonID} [get]
func (h *AvailabilityHandler) ListAvailabilitiesBySalonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	availabilities, err := h.service.ListAvailabilitiesBySalonID(salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(availabilities)
}

// @Summary List availabilities by service ID
// @Description Retrieve all availabilities for a specific service by its service ID
// @Accept  json
// @Produce  json
// @Param serviceID path int true "Service ID"
// @Success 200 {array} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availabilities/service/{serviceID} [get]
func (h *AvailabilityHandler) ListAvailabilitiesByServiceID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID, err := strconv.Atoi(vars["serviceID"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	availabilities, err := h.service.ListAvailabilitiesByServiceID(serviceID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(availabilities)
}

// @Summary List availabilities by status
// @Description Retrieve all availabilities with a specific status
// @Accept  json
// @Produce  json
// @Param status path string true "Status"
// @Success 200 {array} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availabilities/status/{status} [get]
func (h *AvailabilityHandler) ListAvailabilitiesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	availabilities, err := h.service.ListAvailabilitiesByStatus(status)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(availabilities)
}

// @Summary List open availabilities
// @Description Retrieve all open (available) time slots for a specific service and salon
// @Accept  json
// @Produce  json
// @Param serviceID path int true "Service ID"
// @Param salonID path int true "Salon ID"
// @Success 200 {array} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availabilities/open/{serviceID}/{salonID} [get]
func (h *AvailabilityHandler) ListOpenAvailabilities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID, err := strconv.Atoi(vars["serviceID"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	availabilities, err := h.service.ListOpenAvailabilities(serviceID, salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(availabilities)
}

// @Summary Book an availability
// @Description Book an available time slot by its ID and update its status to "Booked"
// @Accept  json
// @Produce  json
// @Param availabilityID path int true "Availability ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availability/{availabilityID}/book [put]
func (h *AvailabilityHandler) BookAvailability(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	availabilityID, err := strconv.Atoi(vars["availabilityID"])
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	err = h.service.BookAvailability(availabilityID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Cancel booking
// @Description Cancel a booked time slot by its ID and update its status to "Open"
// @Accept  json
// @Produce  json
// @Param availabilityID path int true "Availability ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availability/{availabilityID}/cancel [put]
func (h *AvailabilityHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	availabilityID, err := strconv.Atoi(vars["availabilityID"])
	if err != nil {
		http.Error(w, "Invalid availability ID", http.StatusBadRequest)
		return
	}

	err = h.service.CancelBooking(availabilityID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary List booked availabilities by service ID and salon ID
// @Description Retrieve all booked time slots for a specific service and salon
// @Accept  json
// @Produce  json
// @Param serviceID path int true "Service ID"
// @Param salonID path int true "Salon ID"
// @Success 200 {array} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availabilities/booked/{serviceID}/{salonID} [get]
func (h *AvailabilityHandler) ListBookedAvailabilities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID, err := strconv.Atoi(vars["serviceID"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	availabilities, err := h.service.ListBookedAvailabilities(serviceID, salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(availabilities)
}

// @Summary List availabilities by date range
// @Description Retrieve all availabilities between the specified start and end dates
// @Accept  json
// @Produce  json
// @Param startDate query string true "Start Date (RFC3339 format)"
// @Param endDate query string true "End Date (RFC3339 format)"
// @Success 200 {array} models.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /availabilities/range [get]
func (h *AvailabilityHandler) ListAvailabilitiesByDateRange(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	// Check if the startDate and endDate strings are in RFC3339 format
	_, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	_, err = time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	// Now that we've validated the date formats, you can pass startDateStr and endDateStr as strings to the service method.
	availabilities, err := h.service.ListAvailabilitiesByDateRange(startDateStr, endDateStr)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(availabilities)
}

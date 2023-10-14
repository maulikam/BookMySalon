package appointment

import (
	"bookmysalon/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type AppointmentHandler struct {
	service AppointmentService
}

func NewAppointmentHandler(s AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: s}
}

// @Summary Create a new appointment
// @Description Create a new appointment with the input payload
// @Accept  json
// @Produce  json
// @Param appointment body models.Appointment true "Create Appointment"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointment [post]
func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment

	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newAppointment, err := h.service.Create(&appointment)
	if err != nil {
		log.Println("Failed to create appointment:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newAppointment)
}

// @Summary Get appointment details
// @Description Get details of an appointment by ID
// @Accept  json
// @Produce  json
// @Param appointmentID path int true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string "Appointment Not Found"
// @Failure 500 {object} map[string]string
// @Router /appointment/{appointmentID} [get]
func (h *AppointmentHandler) GetAppointmentDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	appointment, err := h.service.GetByID(appointmentID)
	if err != nil {
		switch err {
		case ErrAppointmentNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(appointment)
}

// @Summary Update appointment details
// @Description Update details of an existing appointment
// @Accept  json
// @Produce  json
// @Param appointment body models.Appointment true "Update Appointment"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointment/update [put]
func (h *AppointmentHandler) UpdateAppointmentDetails(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment

	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	updatedAppointment, err := h.service.Update(&appointment)
	if err != nil {
		log.Println("Failed to update appointment:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedAppointment)
}

// @Summary Delete an appointment
// @Description Delete an appointment by its ID
// @Accept  json
// @Produce  json
// @Param appointmentID path int true "Appointment ID"
// @Success 204 "Successfully deleted"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string "Appointment Not Found"
// @Failure 500 {object} map[string]string
// @Router /appointment/{appointmentID} [delete]
func (h *AppointmentHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(appointmentID)
	if err != nil {
		switch err {
		case ErrAppointmentNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary List appointments by user ID
// @Description Retrieve all appointments of a specific user by their user ID
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointments/user/{userID} [get]
func (h *AppointmentHandler) ListAppointmentsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.ListByUserID(userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary List appointments by salon ID
// @Description Retrieve all appointments of a specific salon by its salon ID
// @Accept  json
// @Produce  json
// @Param salonID path int true "Salon ID"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointments/salon/{salonID} [get]
func (h *AppointmentHandler) ListAppointmentsBySalonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.ListBySalonID(salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary List appointments by service ID
// @Description Retrieve all appointments for a specific service by its service ID
// @Accept  json
// @Produce  json
// @Param serviceID path int true "Service ID"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointments/service/{serviceID} [get]
func (h *AppointmentHandler) ListAppointmentsByServiceID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceID, err := strconv.Atoi(vars["serviceID"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.ListByServiceID(serviceID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary List appointments by status
// @Description Retrieve all appointments with a specific status
// @Accept  json
// @Produce  json
// @Param status path string true "Status"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointments/status/{status} [get]
func (h *AppointmentHandler) ListAppointmentsByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	appointments, err := h.service.ListByStatus(status)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary Set notification settings for an appointment
// @Description Update the notification settings of an appointment by its ID
// @Accept  json
// @Produce  json
// @Param appointmentID path int true "Appointment ID"
// @Param notificationSetting body string true "Notification Setting"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointment/{appointmentID}/notification [put]
func (h *AppointmentHandler) SetNotificationForAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var notificationSetting string
	if err := json.NewDecoder(r.Body).Decode(&notificationSetting); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	appointment, err := h.service.SetNotification(appointmentID, notificationSetting)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointment)
}

// @Summary List upcoming appointments
// @Description Retrieve all upcoming appointments for the current day and beyond
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Appointment
// @Failure 500 {object} map[string]string
// @Router /appointments/upcoming [get]
func (h *AppointmentHandler) ListUpcomingAppointments(w http.ResponseWriter, r *http.Request) {
	appointments, err := h.service.ListUpcoming()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary List past appointments
// @Description Retrieve all past appointments
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Appointment
// @Failure 500 {object} map[string]string
// @Router /appointments/past [get]
func (h *AppointmentHandler) ListPastAppointments(w http.ResponseWriter, r *http.Request) {
	appointments, err := h.service.ListPast()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary List appointments by date range
// @Description Retrieve all appointments between the specified start and end dates
// @Accept  json
// @Produce  json
// @Param startDate query string true "Start Date (RFC3339 format)"
// @Param endDate query string true "End Date (RFC3339 format)"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointments/range [get]
func (h *AppointmentHandler) ListAppointmentsByDateRange(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.ListByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

// @Summary Cancel an appointment
// @Description Cancel an appointment by ID and update its status to "Cancelled"
// @Accept  json
// @Produce  json
// @Param appointmentID path int true "Appointment ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointment/{appointmentID}/cancel [put]
func (h *AppointmentHandler) CancelAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	err = h.service.Cancel(appointmentID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Confirm an appointment
// @Description Confirm an appointment by ID and update its status to "Confirmed"
// @Accept  json
// @Produce  json
// @Param appointmentID path int true "Appointment ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointment/{appointmentID}/confirm [put]
func (h *AppointmentHandler) ConfirmAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	err = h.service.Confirm(appointmentID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Reschedule an appointment
// @Description Reschedule an appointment by ID and update its date and time
// @Accept  json
// @Produce  json
// @Param appointmentID path int true "Appointment ID"
// @Param newDateTime body string true "New Date and Time (RFC3339 format)"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointment/{appointmentID}/reschedule [put]
func (h *AppointmentHandler) RescheduleAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var newDateTime struct {
		NewDateTime string `json:"newDateTime"`
	}

	if err := json.NewDecoder(r.Body).Decode(&newDateTime); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	updatedAppointment, err := h.service.Reschedule(appointmentID, newDateTime.NewDateTime)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedAppointment)
}

// @Summary List appointments by notification setting
// @Description Retrieve all appointments with a specific notification setting
// @Accept  json
// @Produce  json
// @Param setting query string true "Notification Setting"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /appointments/notification [get]
func (h *AppointmentHandler) ListAppointmentsByNotificationSetting(w http.ResponseWriter, r *http.Request) {
	setting := r.URL.Query().Get("setting")

	appointments, err := h.service.ListByNotificationSetting(setting)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

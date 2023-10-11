package appointment

import (
	"bookmysalon/models"
	"time"
)

// AppointmentService defines the methods for interacting with the Appointment model.
type AppointmentService interface {
	// Create creates a new appointment and returns the created appointment.
	Create(appointment *models.Appointment) (*models.Appointment, error)

	// GetByID retrieves an appointment based on the given appointment ID.
	GetByID(appointmentID int) (*models.Appointment, error)

	// Update updates an existing appointment based on the given appointment.
	Update(appointment *models.Appointment) (*models.Appointment, error)

	// Delete deletes an appointment based on the given appointment ID.
	Delete(appointmentID int) error

	// ListByUserID retrieves all appointments of a specific user.
	ListByUserID(userID int) ([]*models.Appointment, error)

	// ListBySalonID retrieves all appointments of a specific salon.
	ListBySalonID(salonID int) ([]*models.Appointment, error)

	// ListByServiceID retrieves all appointments for a specific service.
	ListByServiceID(serviceID int) ([]*models.Appointment, error)

	// ListByStatus retrieves all appointments with a specific status.
	ListByStatus(status string) ([]*models.Appointment, error)

	// SetNotification updates the notification settings of an appointment.
	SetNotification(appointmentID int, notificationSetting string) (*models.Appointment, error)

	// ListUpcoming retrieves all upcoming appointments for the current day and beyond.
	ListUpcoming() ([]*models.Appointment, error)

	// ListPast retrieves all past appointments.
	ListPast() ([]*models.Appointment, error)

	// ListByDateRange retrieves all appointments between the specified start and end dates.
	ListByDateRange(startDate, endDate time.Time) ([]*models.Appointment, error)

	// Cancel cancels an appointment and updates its status to "Cancelled".
	Cancel(appointmentID int) error

	// Confirm confirms an appointment and updates its status to "Confirmed".
	Confirm(appointmentID int) error

	// Reschedule changes the date and time of an existing appointment.
	Reschedule(appointmentID int, newDateTime string) (*models.Appointment, error)

	// ListByNotificationSetting retrieves all appointments with a specific notification setting (e.g., "Email" or "SMS").
	ListByNotificationSetting(setting string) ([]*models.Appointment, error)
}

// bookmysalon/models/appointment.go

package models

// Appointment represents a user's booking for a salon service.
// swagger:model
type Appointment struct {
	// The unique ID for the appointment.
	//
	// required: true
	// example: 1
	AppointmentID int `json:"appointment_id"`

	// The ID of the user who booked the appointment.
	//
	// required: true
	// example: 7
	UserID int `json:"user_id"`

	// The ID of the salon where the appointment was booked.
	//
	// required: true
	// example: 5
	SalonID int `json:"salon_id"`

	// The ID of the service booked in this appointment.
	//
	// required: true
	// example: 3
	ServiceID int `json:"service_id"`

	// The date and time of the appointment.
	//
	// required: true
	// example: "2023-07-12T14:00:00Z"
	DateTime string `json:"date_time"`

	// The current status of the appointment (e.g., "Confirmed", "Cancelled").
	//
	// required: true
	// example: "Confirmed"
	Status string `json:"status"`

	// User's notification settings for the appointment (e.g., "Email", "SMS").
	//
	// required: true
	// example: "Email"
	NotificationSettings string `json:"notification_settings"`
}

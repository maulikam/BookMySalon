// bookmysalon/models/appointment.go

package models

type Appointment struct {
	AppointmentID        int    `json:"appointment_id"`
	UserID               int    `json:"user_id"`
	SalonID              int    `json:"salon_id"`
	ServiceID            int    `json:"service_id"`
	DateTime             string `json:"date_time"`
	Status               string `json:"status"`
	NotificationSettings string `json:"notification_settings"`
}

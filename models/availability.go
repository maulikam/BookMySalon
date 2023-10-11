// bookmysalon/models/availability.go

package models

type Availability struct {
	AvailabilityID int    `json:"availability_id"`
	SalonID        int    `json:"salon_id"`
	ServiceID      int    `json:"service_id"`
	StartDateTime  string `json:"start_date_time"`
	EndDateTime    string `json:"end_date_time"`
	Status         string `json:"status"`
}

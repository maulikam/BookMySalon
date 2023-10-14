// bookmysalon/models/availability.go

package models

// Availability represents the available time slots for a salon service.
// swagger:model
type Availability struct {
	// The unique ID for the availability entry.
	//
	// required: true
	// example: 101
	AvailabilityID int `json:"availability_id"`

	// The ID of the salon.
	//
	// required: true
	// example: 5
	SalonID int `json:"salon_id"`

	// The ID of the service provided during this availability time.
	//
	// required: true
	// example: 12
	ServiceID int `json:"service_id"`

	// The starting date and time of the available slot.
	//
	// required: true
	// example: "2023-07-10T10:00:00Z"
	StartDateTime string `json:"start_date_time"`

	// The ending date and time of the available slot.
	//
	// required: true
	// example: "2023-07-10T11:00:00Z"
	EndDateTime string `json:"end_date_time"`

	// The status of the availability (e.g., "Open", "Booked").
	//
	// required: true
	// example: "Open"
	Status string `json:"status"`
}

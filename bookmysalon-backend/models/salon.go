// bookmysalon/models/salon.go

package models

// Salon represents details of a salon in the system.
// swagger:model
type Salon struct {
	// The unique ID for the salon.
	//
	// required: true
	// example: 1
	SalonID int `json:"salon_id"`

	// The name of the salon.
	//
	// required: true
	// example: "Elegant Beauty Parlor"
	Name string `json:"name"`

	// The address of the salon.
	//
	// required: true
	// example: "123 Beauty St, Pleasantville, 12345"
	Address string `json:"address"`

	// The contact details for the salon.
	//
	// required: true
	// example: "(123) 456-7890"
	ContactDetails string `json:"contact_details"`

	// URL of the photos related to the salon.
	//
	// required: false
	// example: "http://example.com/path/to/salon/photo.jpg"
	Photos string `json:"photos"`

	// The average rating for the salon out of 5.
	//
	// required: true
	// example: 4.5
	AverageRating float64 `json:"average_rating"`
}

// Service represents a specific service provided by a salon.
// swagger:model
type Service struct {
	// The unique ID for the service.
	//
	// required: true
	// example: 1
	ServiceID int `json:"service_id"`

	// The ID of the salon offering this service.
	//
	// required: true
	// example: 1
	SalonID int `json:"salon_id"`

	// The name of the service.
	//
	// required: true
	// example: "Haircut"
	Name string `json:"name"`

	// A brief description of the service.
	//
	// required: true
	// example: "A stylish haircut tailored to your preference."
	Description string `json:"description"`

	// The duration of the service.
	//
	// required: true
	// example: "45 minutes"
	Duration string `json:"duration"`

	// The price of the service.
	//
	// required: true
	// example: 25.00
	Price float64 `json:"price"`
}

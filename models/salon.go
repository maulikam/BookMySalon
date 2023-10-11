// bookmysalon/models/salon.go

package models

type Salon struct {
	SalonID        int     `json:"salon_id"`
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	ContactDetails string  `json:"contact_details"`
	Photos         string  `json:"photos"`
	AverageRating  float64 `json:"average_rating"`
}

type Service struct {
	ServiceID   int     `json:"service_id"`
	SalonID     int     `json:"salon_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Duration    string  `json:"duration"`
	Price       float64 `json:"price"`
}

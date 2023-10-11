package salon

import "bookmysalon/models"

// SalonService represents the interface for managing salons
type SalonService interface {
	// Add a new salon and return its ID or an error.
	AddSalon(salon models.Salon) (int, error)

	// Update the details of an existing salon or return an error.
	UpdateSalon(salon models.Salon) error
}

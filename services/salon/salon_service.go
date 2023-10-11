package salon

import "bookmysalon/models"

// SalonService represents the interface for managing salons
type SalonService interface {
	// Add a new salon and return its ID or an error.
	AddSalon(salon models.Salon) (int, error)

	// Update the details of an existing salon or return an error.
	UpdateSalon(salon models.Salon) error

	// Delete a salon by its ID.
	DeleteSalon(salonID int) error

	// Retrieve a salon by its ID.
	GetSalonByID(salonID int) (*models.Salon, error)

	// List all salons in the system.
	ListSalons() ([]models.Salon, error)

	// Add a new service for a salon and return its ID or an error.
	AddService(service models.Service) (int, error)

	// Update the details of a service.
	UpdateService(service models.Service) error

	// Delete a service by its ID.
	DeleteService(serviceID int) error

	// Retrieve a service by its ID.
	GetServiceByID(serviceID int) (*models.Service, error)

	// List all services offered by a specific salon.
	ListServicesBySalon(salonID int) ([]models.Service, error)

	// Get the average rating of a salon.
	GetAverageRating(salonID int) (float64, error)
}

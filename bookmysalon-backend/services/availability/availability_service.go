package availability

import "bookmysalon/models"

// AvailabilityService defines the methods for managing salon service availabilities.
type AvailabilityService interface {
	// CreateAvailability creates a new availability entry.
	CreateAvailability(availability *models.Availability) (*models.Availability, error)

	// GetAvailabilityByID retrieves an availability entry by its unique ID.
	GetAvailabilityByID(availabilityID int) (*models.Availability, error)

	// UpdateAvailability updates the details of an existing availability entry.
	UpdateAvailability(availability *models.Availability) (*models.Availability, error)

	// DeleteAvailability deletes an availability entry by its unique ID.
	DeleteAvailability(availabilityID int) error

	// ListAvailabilitiesBySalonID retrieves all availabilities for a specific salon by its ID.
	ListAvailabilitiesBySalonID(salonID int) ([]*models.Availability, error)

	// ListAvailabilitiesByServiceID retrieves all availabilities for a specific service by its ID.
	ListAvailabilitiesByServiceID(serviceID int) ([]*models.Availability, error)

	// ListAvailabilitiesByStatus retrieves all availabilities with a specific status.
	ListAvailabilitiesByStatus(status string) ([]*models.Availability, error)

	// ListOpenAvailabilities retrieves all open (available) time slots for a specific service and salon.
	ListOpenAvailabilities(serviceID, salonID int) ([]*models.Availability, error)

	// BookAvailability books an available time slot, updating its status to "Booked."
	BookAvailability(availabilityID int) error

	// CancelBooking cancels a booked time slot, updating its status to "Open."
	CancelBooking(availabilityID int) error

	// ListBookedAvailabilities retrieves all booked time slots for a specific service and salon.
	ListBookedAvailabilities(serviceID, salonID int) ([]*models.Availability, error)

	// ListAvailabilitiesByDateRange retrieves all availabilities between the specified start and end dates.
	ListAvailabilitiesByDateRange(startDate, endDate string) ([]*models.Availability, error)
}

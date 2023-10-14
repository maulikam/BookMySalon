package availability

import (
	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"database/sql"
	"errors"
	"log"
)

var ErrAvailabilityNotFound = errors.New("availability not found")

// Constants for error messages.
const (
	ErrorAvailabilityInsert = "Error inserting availability"
	ErrorAvailabilityUpdate = "Error updating availability"
)

// availabilityServiceImpl is the implementation of the AvailabilityService interface.
type availabilityServiceImpl struct {
	db *sql.DB
}

// NewAvailabilityService initializes and returns an instance of AvailabilityService.
func NewAvailabilityService() (AvailabilityService, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &availabilityServiceImpl{
		db: db,
	}, nil
}

// CreateAvailability creates a new availability entry.
func (s *availabilityServiceImpl) CreateAvailability(availability *models.Availability) (*models.Availability, error) {
	const query = `
		INSERT INTO availability(salon_id, service_id, start_date_time, end_date_time, status) 
		VALUES($1, $2, $3, $4, $5) RETURNING availability_id
	`

	var availabilityID int
	err := s.db.QueryRow(query, availability.SalonID, availability.ServiceID, availability.StartDateTime, availability.EndDateTime, availability.Status).Scan(&availabilityID)
	if err != nil {
		log.Printf("%s: %v", ErrorAvailabilityInsert, err)
		return nil, err
	}

	availability.AvailabilityID = availabilityID
	return availability, nil
}

// GetAvailabilityByID retrieves an availability entry by its unique ID.
func (s *availabilityServiceImpl) GetAvailabilityByID(availabilityID int) (*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE availability_id=$1
	`

	var availability models.Availability
	err := s.db.QueryRow(query, availabilityID).Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAvailabilityNotFound
		}
		log.Printf("Error retrieving availability by ID: %v", err)
		return nil, err
	}

	return &availability, nil
}

// UpdateAvailability updates the details of an existing availability entry.
func (s *availabilityServiceImpl) UpdateAvailability(availability *models.Availability) (*models.Availability, error) {
	if availability.AvailabilityID == 0 {
		return nil, errors.New("availability ID must be provided for update")
	}

	const query = `
		UPDATE availability SET salon_id=$1, service_id=$2, start_date_time=$3, end_date_time=$4, status=$5 
		WHERE availability_id=$6
	`

	_, err := s.db.Exec(query, availability.SalonID, availability.ServiceID, availability.StartDateTime, availability.EndDateTime, availability.Status, availability.AvailabilityID)
	if err != nil {
		log.Printf("%s: %v", ErrorAvailabilityUpdate, err)
		return nil, err
	}

	return availability, nil
}

// DeleteAvailability deletes an availability entry by its unique ID.
func (s *availabilityServiceImpl) DeleteAvailability(availabilityID int) error {
	const query = `DELETE FROM availability WHERE availability_id=$1`

	_, err := s.db.Exec(query, availabilityID)
	if err != nil {
		log.Printf("Error deleting availability: %v", err)
		return err
	}

	return nil
}

// ListAvailabilitiesBySalonID retrieves all availabilities for a specific salon by its ID.
func (s *availabilityServiceImpl) ListAvailabilitiesBySalonID(salonID int) ([]*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE salon_id=$1
	`

	rows, err := s.db.Query(query, salonID)
	if err != nil {
		log.Printf("Error listing availabilities by salon: %v", err)
		return nil, err
	}
	defer rows.Close()

	var availabilities []*models.Availability
	for rows.Next() {
		var availability models.Availability
		if err := rows.Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status); err != nil {
			log.Printf("Error scanning availability row: %v", err)
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}

	return availabilities, nil
}

// ListAvailabilitiesByServiceID retrieves all availabilities for a specific service by its ID.
func (s *availabilityServiceImpl) ListAvailabilitiesByServiceID(serviceID int) ([]*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE service_id=$1
	`

	rows, err := s.db.Query(query, serviceID)
	if err != nil {
		log.Printf("Error listing availabilities by service: %v", err)
		return nil, err
	}
	defer rows.Close()

	var availabilities []*models.Availability
	for rows.Next() {
		var availability models.Availability
		if err := rows.Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status); err != nil {
			log.Printf("Error scanning availability row: %v", err)
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}

	return availabilities, nil
}

// ListAvailabilitiesByStatus retrieves all availabilities with a specific status.
func (s *availabilityServiceImpl) ListAvailabilitiesByStatus(status string) ([]*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE status=$1
	`

	rows, err := s.db.Query(query, status)
	if err != nil {
		log.Printf("Error listing availabilities by status: %v", err)
		return nil, err
	}
	defer rows.Close()

	var availabilities []*models.Availability
	for rows.Next() {
		var availability models.Availability
		if err := rows.Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status); err != nil {
			log.Printf("Error scanning availability row: %v", err)
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}

	return availabilities, nil
}

// ListOpenAvailabilities retrieves all open (available) time slots for a specific service and salon.
func (s *availabilityServiceImpl) ListOpenAvailabilities(serviceID, salonID int) ([]*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE salon_id=$1 AND service_id=$2 AND status='Open'
	`

	rows, err := s.db.Query(query, salonID, serviceID)
	if err != nil {
		log.Printf("Error listing open availabilities: %v", err)
		return nil, err
	}
	defer rows.Close()

	var availabilities []*models.Availability
	for rows.Next() {
		var availability models.Availability
		if err := rows.Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status); err != nil {
			log.Printf("Error scanning availability row: %v", err)
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}

	return availabilities, nil
}

// BookAvailability books an available time slot, updating its status to "Booked."
func (s *availabilityServiceImpl) BookAvailability(availabilityID int) error {
	const query = `UPDATE availability SET status='Booked' WHERE availability_id=$1`

	_, err := s.db.Exec(query, availabilityID)
	if err != nil {
		log.Printf("Error booking availability: %v", err)
		return err
	}

	return nil
}

// CancelBooking cancels a booked time slot, updating its status to "Open."
func (s *availabilityServiceImpl) CancelBooking(availabilityID int) error {
	const query = `UPDATE availability SET status='Open' WHERE availability_id=$1`

	_, err := s.db.Exec(query, availabilityID)
	if err != nil {
		log.Printf("Error canceling booking: %v", err)
		return err
	}

	return nil
}

// ListBookedAvailabilities retrieves all booked time slots for a specific service and salon.
func (s *availabilityServiceImpl) ListBookedAvailabilities(serviceID, salonID int) ([]*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE salon_id=$1 AND service_id=$2 AND status='Booked'
	`

	rows, err := s.db.Query(query, salonID, serviceID)
	if err != nil {
		log.Printf("Error listing booked availabilities: %v", err)
		return nil, err
	}
	defer rows.Close()

	var availabilities []*models.Availability
	for rows.Next() {
		var availability models.Availability
		if err := rows.Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status); err != nil {
			log.Printf("Error scanning availability row: %v", err)
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}

	return availabilities, nil
}

// ListAvailabilitiesByDateRange retrieves all availabilities between the specified start and end dates.
func (s *availabilityServiceImpl) ListAvailabilitiesByDateRange(startDate, endDate string) ([]*models.Availability, error) {
	const query = `
		SELECT availability_id, salon_id, service_id, start_date_time, end_date_time, status
		FROM availability WHERE start_date_time >= $1 AND end_date_time <= $2
	`

	rows, err := s.db.Query(query, startDate, endDate)
	if err != nil {
		log.Printf("Error listing availabilities by date range: %v", err)
		return nil, err
	}
	defer rows.Close()

	var availabilities []*models.Availability
	for rows.Next() {
		var availability models.Availability
		if err := rows.Scan(&availability.AvailabilityID, &availability.SalonID, &availability.ServiceID, &availability.StartDateTime, &availability.EndDateTime, &availability.Status); err != nil {
			log.Printf("Error scanning availability row: %v", err)
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}

	return availabilities, nil
}

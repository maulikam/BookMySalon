package salon

import (
	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"database/sql"
	"errors"
	"log"
)

var ErrSalonNotFound = errors.New("salon not found")

// Constants for error messages.
const (
	ErrorSalonInsert   = "Error inserting salon"
	ErrorSalonUpdate   = "Error updating salon"
	ErrorSalonIDNotSet = "salon ID must be provided for update"
)

// salonServiceImpl is the implementation of the SalonService interface.
type salonServiceImpl struct {
	db *sql.DB
}

// NewSalonService initializes and returns an instance of SalonService.
func NewSalonService() (SalonService, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &salonServiceImpl{
		db: db,
	}, nil
}

// AddSalon adds a new salon to the database and returns its ID.
// swagger:model
func (s *salonServiceImpl) AddSalon(salon models.Salon) (int, error) {
	const query = `
		INSERT INTO salons(name, address, contact_details, photos, average_rating) 
		VALUES($1, $2, $3, $4, $5) RETURNING salon_id
	`

	var salonID int
	err := s.db.QueryRow(query, salon.Name, salon.Address, salon.ContactDetails, salon.Photos, salon.AverageRating).Scan(&salonID)
	if err != nil {
		log.Printf("%s: %v", ErrorSalonInsert, err)
		return 0, err
	}

	return salonID, nil
}

// UpdateSalon updates the details of an existing salon.
// swagger:model
func (s *salonServiceImpl) UpdateSalon(salon models.Salon) error {
	if salon.SalonID == 0 {
		return errors.New(ErrorSalonIDNotSet)
	}

	const query = `
		UPDATE salons SET name=$1, address=$2, contact_details=$3, photos=$4, average_rating=$5 
		WHERE salon_id=$6
	`

	_, err := s.db.Exec(query, salon.Name, salon.Address, salon.ContactDetails, salon.Photos, salon.AverageRating, salon.SalonID)
	if err != nil {
		log.Printf("%s: %v", ErrorSalonUpdate, err)
		return err
	}

	return nil
}

// DeleteSalon deletes a salon from the database using its ID.
func (s *salonServiceImpl) DeleteSalon(salonID int) error {
	const query = `DELETE FROM salons WHERE salon_id=$1`

	_, err := s.db.Exec(query, salonID)
	if err != nil {
		log.Printf("Error deleting salon: %v", err)
		return err
	}

	return nil
}

// GetSalonByID retrieves a salon by its ID.
func (s *salonServiceImpl) GetSalonByID(salonID int) (*models.Salon, error) {
	const query = `
		SELECT salon_id, name, address, contact_details, photos, average_rating
		FROM salons WHERE salon_id=$1
	`

	var salon models.Salon
	err := s.db.QueryRow(query, salonID).Scan(&salon.SalonID, &salon.Name, &salon.Address, &salon.ContactDetails, &salon.Photos, &salon.AverageRating)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrSalonNotFound
		}
		log.Printf("Error retrieving salon by ID: %v", err)
		return nil, err
	}

	return &salon, nil
}

// ListSalons retrieves all salons from the database.
func (s *salonServiceImpl) ListSalons() ([]models.Salon, error) {
	const query = `
		SELECT salon_id, name, address, contact_details, photos, average_rating
		FROM salons
	`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("Error listing salons: %v", err)
		return nil, err
	}
	defer rows.Close()

	var salons []models.Salon
	for rows.Next() {
		var salon models.Salon
		if err := rows.Scan(&salon.SalonID, &salon.Name, &salon.Address, &salon.ContactDetails, &salon.Photos, &salon.AverageRating); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		salons = append(salons, salon)
	}

	return salons, nil
}

// AddService adds a new service to the database and returns its ID.
func (s *salonServiceImpl) AddService(service models.Service) (int, error) {
	const query = `
		INSERT INTO services(salon_id, name, description, duration, price) 
		VALUES($1, $2, $3, $4, $5) RETURNING service_id
	`

	var serviceID int
	err := s.db.QueryRow(query, service.SalonID, service.Name, service.Description, service.Duration, service.Price).Scan(&serviceID)
	if err != nil {
		log.Printf("Error inserting service: %v", err)
		return 0, err
	}

	return serviceID, nil
}

// UpdateService updates the details of a service in the database.
func (s *salonServiceImpl) UpdateService(service models.Service) error {
	const query = `
		UPDATE services SET salon_id=$1, name=$2, description=$3, duration=$4, price=$5 
		WHERE service_id=$6
	`

	_, err := s.db.Exec(query, service.SalonID, service.Name, service.Description, service.Duration, service.Price, service.ServiceID)
	if err != nil {
		log.Printf("Error updating service: %v", err)
		return err
	}

	return nil
}

// DeleteService deletes a service from the database using its ID.
func (s *salonServiceImpl) DeleteService(serviceID int) error {
	const query = `DELETE FROM services WHERE service_id=$1`

	_, err := s.db.Exec(query, serviceID)
	if err != nil {
		log.Printf("Error deleting service: %v", err)
		return err
	}

	return nil
}

// GetServiceByID retrieves a service by its ID.
func (s *salonServiceImpl) GetServiceByID(serviceID int) (*models.Service, error) {
	const query = `
		SELECT service_id, salon_id, name, description, duration, price
		FROM services WHERE service_id=$1
	`

	var service models.Service
	err := s.db.QueryRow(query, serviceID).Scan(&service.ServiceID, &service.SalonID, &service.Name, &service.Description, &service.Duration, &service.Price)
	if err != nil {
		log.Printf("Error retrieving service by ID: %v", err)
		return nil, err
	}

	return &service, nil
}

// ListServicesBySalon retrieves all services offered by a specific salon.
func (s *salonServiceImpl) ListServicesBySalon(salonID int) ([]models.Service, error) {
	const query = `
		SELECT service_id, salon_id, name, description, duration, price
		FROM services WHERE salon_id=$1
	`

	rows, err := s.db.Query(query, salonID)
	if err != nil {
		log.Printf("Error listing services by salon: %v", err)
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		if err := rows.Scan(&service.ServiceID, &service.SalonID, &service.Name, &service.Description, &service.Duration, &service.Price); err != nil {
			log.Printf("Error scanning service row: %v", err)
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

// GetAverageRating retrieves the average rating of a salon.
func (s *salonServiceImpl) GetAverageRating(salonID int) (float64, error) {
	const query = `
		SELECT average_rating
		FROM salons WHERE salon_id=$1
	`

	var avgRating float64
	err := s.db.QueryRow(query, salonID).Scan(&avgRating)
	if err != nil {
		log.Printf("Error retrieving average rating of salon: %v", err)
		return 0, err
	}

	return avgRating, nil
}

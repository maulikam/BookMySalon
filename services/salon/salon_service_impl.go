package salon

import (
	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"database/sql"
	"errors"
	"log"
)

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

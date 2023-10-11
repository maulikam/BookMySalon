package appointment

import (
	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"database/sql"
	"errors"
	"log"
	"time"
)

var ErrAppointmentNotFound = errors.New("appointment not found")

const (
	ErrorAppointmentInsert   = "Error inserting appointment"
	ErrorAppointmentUpdate   = "Error updating appointment"
	ErrorAppointmentIDNotSet = "Appointment ID must be provided for update"
	ErrorAppointmentDelete   = "Error deleting appointment"
)

type appointmentServiceImpl struct {
	db *sql.DB
}

// NewAppointmentService initializes and returns an instance of AppointmentService.
func NewAppointmentService() (AppointmentService, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &appointmentServiceImpl{
		db: db,
	}, nil
}

// Create inserts a new appointment into the database.
func (a *appointmentServiceImpl) Create(appointment *models.Appointment) (*models.Appointment, error) {
	const query = `
		INSERT INTO appointments(user_id, salon_id, service_id, date_time, status, notification_settings) 
		VALUES($1, $2, $3, $4, $5, $6) RETURNING appointment_id
	`

	err := a.db.QueryRow(query, appointment.UserID, appointment.SalonID, appointment.ServiceID, appointment.DateTime, appointment.Status, appointment.NotificationSettings).Scan(&appointment.AppointmentID)
	if err != nil {
		log.Printf("%s: %v", ErrorAppointmentInsert, err)
		return nil, err
	}

	return appointment, nil
}

// GetByID retrieves an appointment by its ID.
func (a *appointmentServiceImpl) GetByID(appointmentID int) (*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE appointment_id=$1
	`

	appointment := &models.Appointment{}
	err := a.db.QueryRow(query, appointmentID).Scan(&appointment.AppointmentID, &appointment.UserID, &appointment.SalonID, &appointment.ServiceID, &appointment.DateTime, &appointment.Status, &appointment.NotificationSettings)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAppointmentNotFound
		}
		log.Printf("Error retrieving appointment by ID: %v", err)
		return nil, err
	}

	return appointment, nil
}

// Update modifies the details of an existing appointment.
func (a *appointmentServiceImpl) Update(appointment *models.Appointment) (*models.Appointment, error) {
	if appointment.AppointmentID == 0 {
		return nil, errors.New(ErrorAppointmentIDNotSet)
	}

	const query = `
		UPDATE appointments SET user_id=$1, salon_id=$2, service_id=$3, date_time=$4, status=$5, notification_settings=$6
		WHERE appointment_id=$7
	`

	_, err := a.db.Exec(query, appointment.UserID, appointment.SalonID, appointment.ServiceID, appointment.DateTime, appointment.Status, appointment.NotificationSettings, appointment.AppointmentID)
	if err != nil {
		log.Printf("%s: %v", ErrorAppointmentUpdate, err)
		return nil, err
	}

	return appointment, nil
}

// Delete removes an appointment based on the given appointment ID.
func (a *appointmentServiceImpl) Delete(appointmentID int) error {
	const query = `DELETE FROM appointments WHERE appointment_id=$1`

	_, err := a.db.Exec(query, appointmentID)
	if err != nil {
		log.Printf("%s: %v", ErrorAppointmentDelete, err)
		return err
	}

	return nil
}

// ListByUserID retrieves all appointments of a specific user.
func (a *appointmentServiceImpl) ListByUserID(userID int) ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE user_id=$1
	`
	return a.listByQuery(query, userID)
}

// ListBySalonID retrieves all appointments of a specific salon.
func (a *appointmentServiceImpl) ListBySalonID(salonID int) ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE salon_id=$1
	`
	return a.listByQuery(query, salonID)
}

// ListByServiceID retrieves all appointments for a specific service.
func (a *appointmentServiceImpl) ListByServiceID(serviceID int) ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE service_id=$1
	`
	return a.listByQuery(query, serviceID)
}

// ListByStatus retrieves all appointments with a specific status.
func (a *appointmentServiceImpl) ListByStatus(status string) ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE status=$1
	`
	return a.listByQuery(query, status)
}

// SetNotification updates the notification settings of an appointment.
func (a *appointmentServiceImpl) SetNotification(appointmentID int, notificationSetting string) (*models.Appointment, error) {
	const query = `
		UPDATE appointments SET notification_settings=$1 WHERE appointment_id=$2 RETURNING appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
	`

	appointment := &models.Appointment{}
	err := a.db.QueryRow(query, notificationSetting, appointmentID).Scan(&appointment.AppointmentID, &appointment.UserID, &appointment.SalonID, &appointment.ServiceID, &appointment.DateTime, &appointment.Status, &appointment.NotificationSettings)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

// ListUpcoming retrieves all upcoming appointments for the current day and beyond.
func (a *appointmentServiceImpl) ListUpcoming() ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE date_time >= $1
	`
	return a.listByQuery(query, time.Now())
}

// ListPast retrieves all past appointments.
func (a *appointmentServiceImpl) ListPast() ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE date_time < $1
	`
	return a.listByQuery(query, time.Now())
}

// ListByDateRange retrieves all appointments between the specified start and end dates.
func (a *appointmentServiceImpl) ListByDateRange(startDate, endDate time.Time) ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE date_time BETWEEN $1 AND $2
	`
	return a.listByQueryWithRange(query, startDate, endDate)
}

// listByQueryWithRange is a helper function that executes a given query with a date range and returns a list of appointments.
func (a *appointmentServiceImpl) listByQueryWithRange(query string, startDate, endDate time.Time) ([]*models.Appointment, error) {
	rows, err := a.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*models.Appointment
	for rows.Next() {
		appointment := &models.Appointment{}
		if err := rows.Scan(&appointment.AppointmentID, &appointment.UserID, &appointment.SalonID, &appointment.ServiceID, &appointment.DateTime, &appointment.Status, &appointment.NotificationSettings); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return appointments, nil
}

// listByQuery is a helper function that executes a given query and parameter, and returns a list of appointments.
func (a *appointmentServiceImpl) listByQuery(query string, param interface{}) ([]*models.Appointment, error) {
	rows, err := a.db.Query(query, param)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*models.Appointment
	for rows.Next() {
		appointment := &models.Appointment{}
		if err := rows.Scan(&appointment.AppointmentID, &appointment.UserID, &appointment.SalonID, &appointment.ServiceID, &appointment.DateTime, &appointment.Status, &appointment.NotificationSettings); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return appointments, nil
}

// Cancel cancels an appointment and updates its status to "Cancelled".
func (a *appointmentServiceImpl) Cancel(appointmentID int) error {
	const query = `UPDATE appointments SET status='Cancelled' WHERE appointment_id=$1`

	_, err := a.db.Exec(query, appointmentID)
	if err != nil {
		return err
	}

	return nil
}

// Confirm confirms an appointment and updates its status to "Confirmed".
func (a *appointmentServiceImpl) Confirm(appointmentID int) error {
	const query = `UPDATE appointments SET status='Confirmed' WHERE appointment_id=$1`

	_, err := a.db.Exec(query, appointmentID)
	if err != nil {
		return err
	}

	return nil
}

// Reschedule changes the date and time of an existing appointment.
func (a *appointmentServiceImpl) Reschedule(appointmentID int, newDateTime string) (*models.Appointment, error) {
	const query = `
		UPDATE appointments SET date_time=$1 WHERE appointment_id=$2 RETURNING appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
	`

	appointment := &models.Appointment{}
	err := a.db.QueryRow(query, newDateTime, appointmentID).Scan(&appointment.AppointmentID, &appointment.UserID, &appointment.SalonID, &appointment.ServiceID, &appointment.DateTime, &appointment.Status, &appointment.NotificationSettings)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

// ListByNotificationSetting retrieves all appointments with a specific notification setting (e.g., "Email" or "SMS").
func (a *appointmentServiceImpl) ListByNotificationSetting(setting string) ([]*models.Appointment, error) {
	const query = `
		SELECT appointment_id, user_id, salon_id, service_id, date_time, status, notification_settings
		FROM appointments WHERE notification_settings=$1
	`
	return a.listByQuery(query, setting)
}

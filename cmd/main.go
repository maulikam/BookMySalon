package main

import (
	"bookmysalon/pkg/database"
	"bookmysalon/services/appointment"
	"bookmysalon/services/availability"
	"bookmysalon/services/review"
	"bookmysalon/services/salon"
	"bookmysalon/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const ServerAddress = ":8080"

func main() {

	// Run the migrations first
	database.RunMigrations()

	// Initialize Salon service
	salonService, err := salon.NewSalonService()
	if err != nil {
		log.Fatalf("Failed to initialize salon service: %v", err)
	}
	salonHandler := salon.NewSalonHandler(salonService)

	// Initialize User service
	userServiceImpl := &user.UserServiceImpl{} // Assuming you have implemented it
	userHandler := user.NewUserHandler(userServiceImpl)

	// Initialize Appointment service
	appointmentService, err := appointment.NewAppointmentService()
	if err != nil {
		log.Fatalf("Failed to initialize appointment service: %v", err)
	}
	appointmentHandler := appointment.NewAppointmentHandler(appointmentService)

	// Initialize Availability service
	availabilityService, err := availability.NewAvailabilityService()
	if err != nil {
		log.Fatalf("Failed to initialize availability service: %v", err)
	}
	availabilityHandler := availability.NewAvailabilityHandler(availabilityService)

	// Initialize Review service
	// Initialize Appointment service
	reviewService, err := review.NewReviewService()
	if err != nil {
		log.Fatalf("Failed to initialize appointment service: %v", err)
	}
	reviewHandler := review.NewReviewHandler(reviewService)

	r := mux.NewRouter()

	// Salon routes
	r.HandleFunc("/salon", salonHandler.CreateSalon).Methods("POST")
	r.HandleFunc("/salon/update", salonHandler.UpdateSalonDetails).Methods("PUT")
	r.HandleFunc("/salons", salonHandler.ListAllSalons).Methods("GET")
	r.HandleFunc("/service", salonHandler.AddService).Methods("POST")
	r.HandleFunc("/service/update", salonHandler.UpdateServiceDetails).Methods("PUT")
	r.HandleFunc("/service/{serviceID}", salonHandler.DeleteService).Methods("DELETE")
	r.HandleFunc("/service/{serviceID}", salonHandler.GetServiceDetails).Methods("GET")
	r.HandleFunc("/salon/{salonID}/services", salonHandler.GetServicesBySalon).Methods("GET")
	r.HandleFunc("/salon/{salonID}/average-rating", salonHandler.GetSalonAverageRating).Methods("GET")
	r.HandleFunc("/salon/{salonID}", salonHandler.GetSalonDetails).Methods("GET")
	r.HandleFunc("/salon/{salonID}", salonHandler.DeleteSalon).Methods("DELETE")

	// User routes
	r.HandleFunc("/register", userHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", userHandler.ProfileHandler).Methods("GET")
	r.HandleFunc("/profile", userHandler.UpdateProfileHandler).Methods("PUT")
	r.HandleFunc("/change-password", userHandler.ChangePasswordHandler).Methods("PUT")
	r.HandleFunc("/profile", userHandler.DeleteAccountHandler).Methods("DELETE")

	// Appointment routes
	r.HandleFunc("/appointment", appointmentHandler.CreateAppointment).Methods("POST")
	r.HandleFunc("/appointment/{appointmentID}", appointmentHandler.GetAppointmentDetails).Methods("GET")
	r.HandleFunc("/appointment/update", appointmentHandler.UpdateAppointmentDetails).Methods("PUT")
	r.HandleFunc("/appointment/{appointmentID}", appointmentHandler.DeleteAppointment).Methods("DELETE")
	r.HandleFunc("/appointments/user/{userID}", appointmentHandler.ListAppointmentsByUserID).Methods("GET")
	r.HandleFunc("/appointments/salon/{salonID}", appointmentHandler.ListAppointmentsBySalonID).Methods("GET")
	r.HandleFunc("/appointments/service/{serviceID}", appointmentHandler.ListAppointmentsByServiceID).Methods("GET")
	r.HandleFunc("/appointments/status/{status}", appointmentHandler.ListAppointmentsByStatus).Methods("GET")
	r.HandleFunc("/appointment/{appointmentID}/notification", appointmentHandler.SetNotificationForAppointment).Methods("PUT")
	r.HandleFunc("/appointments/upcoming", appointmentHandler.ListUpcomingAppointments).Methods("GET")
	r.HandleFunc("/appointments/past", appointmentHandler.ListPastAppointments).Methods("GET")
	r.HandleFunc("/appointments/range", appointmentHandler.ListAppointmentsByDateRange).Methods("GET")
	r.HandleFunc("/appointment/{appointmentID}/cancel", appointmentHandler.CancelAppointment).Methods("PUT")
	r.HandleFunc("/appointment/{appointmentID}/confirm", appointmentHandler.ConfirmAppointment).Methods("PUT")
	r.HandleFunc("/appointment/{appointmentID}/reschedule", appointmentHandler.RescheduleAppointment).Methods("PUT")
	r.HandleFunc("/appointments/notification", appointmentHandler.ListAppointmentsByNotificationSetting).Methods("GET")

	// Availability routes
	r.HandleFunc("/availability", availabilityHandler.CreateAvailability).Methods("POST")
	r.HandleFunc("/availability/{availabilityID}", availabilityHandler.GetAvailabilityDetails).Methods("GET")
	r.HandleFunc("/availability/update", availabilityHandler.UpdateAvailabilityDetails).Methods("PUT")
	r.HandleFunc("/availability/{availabilityID}", availabilityHandler.DeleteAvailability).Methods("DELETE")
	r.HandleFunc("/availabilities/salon/{salonID}", availabilityHandler.ListAvailabilitiesBySalonID).Methods("GET")
	r.HandleFunc("/availabilities/service/{serviceID}", availabilityHandler.ListAvailabilitiesByServiceID).Methods("GET")
	r.HandleFunc("/availabilities/status/{status}", availabilityHandler.ListAvailabilitiesByStatus).Methods("GET")
	r.HandleFunc("/availabilities/open/{serviceID}/{salonID}", availabilityHandler.ListOpenAvailabilities).Methods("GET")
	r.HandleFunc("/availability/{availabilityID}/book", availabilityHandler.BookAvailability).Methods("PUT")
	r.HandleFunc("/availability/{availabilityID}/cancel", availabilityHandler.CancelBooking).Methods("PUT")
	r.HandleFunc("/availabilities/booked/{serviceID}/{salonID}", availabilityHandler.ListBookedAvailabilities).Methods("GET")
	r.HandleFunc("/availabilities/range", availabilityHandler.ListAvailabilitiesByDateRange).Methods("GET")

	// Define your review routes
	r.HandleFunc("/reviews", reviewHandler.CreateReview).Methods("POST")
	r.HandleFunc("/reviews/{reviewID}", reviewHandler.GetReviewByID).Methods("GET")
	r.HandleFunc("/reviews/{reviewID}", reviewHandler.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{reviewID}", reviewHandler.DeleteReview).Methods("DELETE")
	r.HandleFunc("/reviews/salon/{salonID}", reviewHandler.ListReviewsBySalonID).Methods("GET")
	r.HandleFunc("/reviews/user/{userID}", reviewHandler.ListReviewsByUserID).Methods("GET")
	r.HandleFunc("/reviews/rating/{rating}", reviewHandler.ListReviewsByRating).Methods("GET")

	// Swagger UI and JSON routes (assuming you still want these from the salon handlers)
	fs := http.FileServer(http.Dir("./swaggerui/"))
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
	r.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger.json")
	})

	log.Printf("Server started on %s", ServerAddress)
	if err := http.ListenAndServe(ServerAddress, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

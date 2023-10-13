package main

import (
	"bookmysalon/pkg/database"
	"bookmysalon/pkg/middleware"
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

// handleInitializationError checks if there's an error and logs it.
func handleInitializationError(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {

	// Run the migrations first
	database.RunMigrations()

	// Services and handlers initialization
	salonService, err := salon.NewSalonService()
	handleInitializationError(err, "Failed to initialize salon service: %v")
	salonHandler := salon.NewSalonHandler(salonService)

	userServiceImpl := &user.UserServiceImpl{}
	userHandler := user.NewUserHandler(userServiceImpl)

	appointmentService, err := appointment.NewAppointmentService()
	handleInitializationError(err, "Failed to initialize appointment service: %v")
	appointmentHandler := appointment.NewAppointmentHandler(appointmentService)

	availabilityService, err := availability.NewAvailabilityService()
	handleInitializationError(err, "Failed to initialize availability service: %v")
	availabilityHandler := availability.NewAvailabilityHandler(availabilityService)

	reviewService, err := review.NewReviewService()
	handleInitializationError(err, "Failed to initialize review service: %v")
	reviewHandler := review.NewReviewHandler(reviewService)

	r := mux.NewRouter()

	// General routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to BookMySalon API!"))
	}).Methods("GET")

	// Salon routes
	r.HandleFunc("/salon", middleware.Authenticate(salonHandler.CreateSalon)).Methods("POST")
	r.HandleFunc("/salon/update", middleware.Authenticate(salonHandler.UpdateSalonDetails)).Methods("PUT")
	r.HandleFunc("/salons", middleware.Authenticate(salonHandler.ListAllSalons)).Methods("GET")
	r.HandleFunc("/service", middleware.Authenticate(salonHandler.AddService)).Methods("POST")
	r.HandleFunc("/service/update", middleware.Authenticate(salonHandler.UpdateServiceDetails)).Methods("PUT")
	r.HandleFunc("/service/{serviceID}", middleware.Authenticate(salonHandler.DeleteService)).Methods("DELETE")
	r.HandleFunc("/service/{serviceID}", middleware.Authenticate(salonHandler.GetServiceDetails)).Methods("GET")
	r.HandleFunc("/salon/{salonID}/services", middleware.Authenticate(salonHandler.GetServicesBySalon)).Methods("GET")
	r.HandleFunc("/salon/{salonID}/average-rating", middleware.Authenticate(salonHandler.GetSalonAverageRating)).Methods("GET")
	r.HandleFunc("/salon/{salonID}", middleware.Authenticate(salonHandler.GetSalonDetails)).Methods("GET")
	r.HandleFunc("/salon/{salonID}", middleware.Authenticate(salonHandler.DeleteSalon)).Methods("DELETE")

	// User routes
	r.HandleFunc("/register", userHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", middleware.Authenticate(userHandler.ProfileHandler)).Methods("GET")
	r.HandleFunc("/profile", middleware.Authenticate(userHandler.UpdateProfileHandler)).Methods("PUT")
	r.HandleFunc("/change-password", middleware.Authenticate(userHandler.ChangePasswordHandler)).Methods("PUT")
	r.HandleFunc("/profile", middleware.Authenticate(userHandler.DeleteAccountHandler)).Methods("DELETE")

	// Appointment routes
	r.HandleFunc("/appointment", middleware.Authenticate(appointmentHandler.CreateAppointment)).Methods("POST")
	r.HandleFunc("/appointment/{appointmentID}", middleware.Authenticate(appointmentHandler.GetAppointmentDetails)).Methods("GET")
	r.HandleFunc("/appointment/update", middleware.Authenticate(appointmentHandler.UpdateAppointmentDetails)).Methods("PUT")
	r.HandleFunc("/appointment/{appointmentID}", middleware.Authenticate(appointmentHandler.DeleteAppointment)).Methods("DELETE")
	r.HandleFunc("/appointments/user/{userID}", middleware.Authenticate(appointmentHandler.ListAppointmentsByUserID)).Methods("GET")
	r.HandleFunc("/appointments/salon/{salonID}", middleware.Authenticate(appointmentHandler.ListAppointmentsBySalonID)).Methods("GET")
	r.HandleFunc("/appointments/service/{serviceID}", middleware.Authenticate(appointmentHandler.ListAppointmentsByServiceID)).Methods("GET")
	r.HandleFunc("/appointments/status/{status}", middleware.Authenticate(appointmentHandler.ListAppointmentsByStatus)).Methods("GET")
	r.HandleFunc("/appointment/{appointmentID}/notification", middleware.Authenticate(appointmentHandler.SetNotificationForAppointment)).Methods("PUT")
	r.HandleFunc("/appointments/upcoming", middleware.Authenticate(appointmentHandler.ListUpcomingAppointments)).Methods("GET")
	r.HandleFunc("/appointments/past", middleware.Authenticate(appointmentHandler.ListPastAppointments)).Methods("GET")
	r.HandleFunc("/appointments/range", middleware.Authenticate(appointmentHandler.ListAppointmentsByDateRange)).Methods("GET")
	r.HandleFunc("/appointment/{appointmentID}/cancel", middleware.Authenticate(appointmentHandler.CancelAppointment)).Methods("PUT")
	r.HandleFunc("/appointment/{appointmentID}/confirm", middleware.Authenticate(appointmentHandler.ConfirmAppointment)).Methods("PUT")
	r.HandleFunc("/appointment/{appointmentID}/reschedule", middleware.Authenticate(appointmentHandler.RescheduleAppointment)).Methods("PUT")
	r.HandleFunc("/appointments/notification", middleware.Authenticate(appointmentHandler.ListAppointmentsByNotificationSetting)).Methods("GET")

	// Availability routes
	r.HandleFunc("/availability", middleware.Authenticate(availabilityHandler.CreateAvailability)).Methods("POST")
	r.HandleFunc("/availability/{availabilityID}", middleware.Authenticate(availabilityHandler.GetAvailabilityDetails)).Methods("GET")
	r.HandleFunc("/availability/update", middleware.Authenticate(availabilityHandler.UpdateAvailabilityDetails)).Methods("PUT")
	r.HandleFunc("/availability/{availabilityID}", middleware.Authenticate(availabilityHandler.DeleteAvailability)).Methods("DELETE")
	r.HandleFunc("/availabilities/salon/{salonID}", middleware.Authenticate(availabilityHandler.ListAvailabilitiesBySalonID)).Methods("GET")
	r.HandleFunc("/availabilities/service/{serviceID}", middleware.Authenticate(availabilityHandler.ListAvailabilitiesByServiceID)).Methods("GET")
	r.HandleFunc("/availabilities/status/{status}", middleware.Authenticate(availabilityHandler.ListAvailabilitiesByStatus)).Methods("GET")
	r.HandleFunc("/availabilities/open/{serviceID}/{salonID}", middleware.Authenticate(availabilityHandler.ListOpenAvailabilities)).Methods("GET")
	r.HandleFunc("/availability/{availabilityID}/book", middleware.Authenticate(availabilityHandler.BookAvailability)).Methods("PUT")
	r.HandleFunc("/availability/{availabilityID}/cancel", middleware.Authenticate(availabilityHandler.CancelBooking)).Methods("PUT")
	r.HandleFunc("/availabilities/booked/{serviceID}/{salonID}", middleware.Authenticate(availabilityHandler.ListBookedAvailabilities)).Methods("GET")
	r.HandleFunc("/availabilities/range", middleware.Authenticate(availabilityHandler.ListAvailabilitiesByDateRange)).Methods("GET")

	// Define your review routes
	r.HandleFunc("/reviews", middleware.Authenticate(reviewHandler.CreateReview)).Methods("POST")
	r.HandleFunc("/reviews/{reviewID}", middleware.Authenticate(reviewHandler.GetReviewByID)).Methods("GET")
	r.HandleFunc("/reviews/{reviewID}", middleware.Authenticate(reviewHandler.UpdateReview)).Methods("PUT")
	r.HandleFunc("/reviews/{reviewID}", middleware.Authenticate(reviewHandler.DeleteReview)).Methods("DELETE")
	r.HandleFunc("/reviews/salon/{salonID}", middleware.Authenticate(reviewHandler.ListReviewsBySalonID)).Methods("GET")
	r.HandleFunc("/reviews/user/{userID}", middleware.Authenticate(reviewHandler.ListReviewsByUserID)).Methods("GET")
	r.HandleFunc("/reviews/rating/{rating}", middleware.Authenticate(reviewHandler.ListReviewsByRating)).Methods("GET")

	// Swagger UI and JSON routes
	fs := http.FileServer(http.Dir("swaggerui"))
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
	r.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger.json")
	})

	log.Printf("Server started on %s", ServerAddress)
	if err := http.ListenAndServe(ServerAddress, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

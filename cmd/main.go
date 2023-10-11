package main

import (
	"bookmysalon/services/appointment"
	"bookmysalon/services/salon"
	"bookmysalon/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const ServerAddress = ":8080"

func main() {
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

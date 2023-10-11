package main

import (
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

	r := mux.NewRouter()

	// Salon routes
	r.HandleFunc("/salon", salonHandler.CreateSalon).Methods("POST")
	r.HandleFunc("/salon/update", salonHandler.UpdateSalonDetails).Methods("PUT")

	// User routes
	r.HandleFunc("/register", userHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", userHandler.ProfileHandler).Methods("GET")
	r.HandleFunc("/profile", userHandler.UpdateProfileHandler).Methods("PUT")
	r.HandleFunc("/change-password", userHandler.ChangePasswordHandler).Methods("PUT")
	r.HandleFunc("/profile", userHandler.DeleteAccountHandler).Methods("DELETE")

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

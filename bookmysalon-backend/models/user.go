// bookmysalon/models/user.go

package models

// User represents the data for a user in the system.
// swagger:model
type User struct {
	// The unique ID for the user.
	//
	// required: true
	// example: 1
	ID int `json:"id"`

	// The username for the user. This needs to be unique across the platform.
	//
	// required: true
	// example: "john_doe"
	Username string `json:"username"`

	// The password for the user. This will be stored encrypted and is never returned in API responses.
	//
	// required: true
	// example: "securepassword123"
	Password string `json:"password"`

	// The email address associated with the user.
	//
	// required: false
	// example: "johndoe@example.com"
	Email string `json:"email"`

	// The profile image URL for the user.
	//
	// required: false
	// example: "http://example.com/path/to/image.jpg"
	ProfileImage string `json:"profile_image"`

	// The date the user joined the platform.
	//
	// required: false
	// example: "2023-01-01"
	DateJoined string `json:"date_joined"`

	// The date of the user's last login.
	//
	// required: false
	// example: "2023-01-10"
	LastLogin string `json:"last_login"`
}

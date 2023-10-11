// bookmysalon/models/user.go

package models

// User representation
// swagger:parameters registerUser loginUser
type User struct {
	// User's unique ID
	//
	// required: true
	ID int `json:"id"`

	// Username for the user
	//
	// required: true
	Username string `json:"username"`

	// Password for the user
	//
	// required: true
	Password     string `json:"password"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
	DateJoined   string `json:"date_joined"`
	LastLogin    string `json:"last_login"`
}

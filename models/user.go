// bookmysalon/models/user.go

package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"` // This should store the hashed value
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
	DateJoined   string `json:"date_joined"`
	LastLogin    string `json:"last_login"`
}

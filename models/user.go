package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Ideally this should be hashed, not plain text.
}

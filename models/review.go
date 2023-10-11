// bookmysalon/models/review.go

package models

type Review struct {
	ReviewID   int    `json:"review_id"`
	UserID     int    `json:"user_id"`
	SalonID    int    `json:"salon_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
}

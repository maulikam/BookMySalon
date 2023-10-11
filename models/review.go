// bookmysalon/models/review.go

package models

// Review represents user reviews for a salon in the system.
// swagger:model
type Review struct {
	// The unique ID for the review.
	//
	// required: true
	// example: 1
	ReviewID int `json:"review_id"`

	// The ID of the user who posted the review.
	//
	// required: true
	// example: 1
	UserID int `json:"user_id"`

	// The ID of the salon for which the review was posted.
	//
	// required: true
	// example: 2
	SalonID int `json:"salon_id"`

	// The rating given by the user, typically out of 5.
	//
	// required: true
	// example: 4
	Rating int `json:"rating"`

	// The comment or feedback given by the user.
	//
	// required: true
	// example: "Great service and friendly staff!"
	Comment string `json:"comment"`

	// The date when the review was posted.
	//
	// required: true
	// example: "2023-05-20"
	DatePosted string `json:"date_posted"`
}

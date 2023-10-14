package review

import (
	"bookmysalon/models"
)

// ReviewService represents the interface for managing user reviews.
type ReviewService interface {
	// CreateReview creates a new review.
	CreateReview(review *models.Review) (*models.Review, error)

	// GetReviewByID retrieves a review by its unique ID.
	GetReviewByID(reviewID int) (*models.Review, error)

	// UpdateReview updates the details of an existing review.
	UpdateReview(review *models.Review) (*models.Review, error)

	// DeleteReview deletes a review by its unique ID.
	DeleteReview(reviewID int) error

	// ListReviewsBySalonID retrieves all reviews for a specific salon by salon ID.
	ListReviewsBySalonID(salonID int) ([]*models.Review, error)

	// ListReviewsByUserID retrieves all reviews posted by a specific user by user ID.
	ListReviewsByUserID(userID int) ([]*models.Review, error)

	// ListReviewsByRating retrieves all reviews with a specific rating.
	ListReviewsByRating(rating int) ([]*models.Review, error)
}

package review

import (
	"bookmysalon/models"
	"bookmysalon/pkg/database"
	"database/sql"
	"errors"
	"log"
)

var (
	ErrReviewNotFound     = errors.New("review not found")
	ErrReviewIDNotSet     = errors.New("review ID must be provided for update")
	ErrorReviewInsert     = "Error inserting review"
	ErrorReviewUpdate     = "Error updating review"
	ErrorReviewDelete     = "Error deleting review"
	ErrorReviewListSalon  = "Error listing reviews by salon ID"
	ErrorReviewListUser   = "Error listing reviews by user ID"
	ErrorReviewListRating = "Error listing reviews by rating"
)

type reviewServiceImpl struct {
	db *sql.DB
}

// NewReviewService initializes and returns an instance of ReviewService.
func NewReviewService() (ReviewService, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &reviewServiceImpl{
		db: db,
	}, nil
}

func (s *reviewServiceImpl) CreateReview(review *models.Review) (*models.Review, error) {
	const query = `
		INSERT INTO reviews(user_id, salon_id, rating, comment, date_posted)
		VALUES($1, $2, $3, $4, $5) RETURNING review_id
	`

	var reviewID int
	err := s.db.QueryRow(query, review.UserID, review.SalonID, review.Rating, review.Comment, review.DatePosted).Scan(&reviewID)
	if err != nil {
		log.Printf("%s: %v", ErrorReviewInsert, err)
		return nil, err
	}

	review.ReviewID = reviewID
	return review, nil
}

func (s *reviewServiceImpl) GetReviewByID(reviewID int) (*models.Review, error) {
	const query = `
		SELECT review_id, user_id, salon_id, rating, comment, date_posted
		FROM reviews WHERE review_id = $1
	`

	var review models.Review
	err := s.db.QueryRow(query, reviewID).Scan(
		&review.ReviewID, &review.UserID, &review.SalonID, &review.Rating, &review.Comment, &review.DatePosted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrReviewNotFound
		}
		log.Printf("Error retrieving review by ID: %v", err)
		return nil, err
	}

	return &review, nil
}

func (s *reviewServiceImpl) UpdateReview(review *models.Review) (*models.Review, error) {
	if review.ReviewID == 0 {
		return nil, ErrReviewIDNotSet
	}

	const query = `
		UPDATE reviews SET user_id=$1, salon_id=$2, rating=$3, comment=$4, date_posted=$5
		WHERE review_id=$6
	`

	_, err := s.db.Exec(query, review.UserID, review.SalonID, review.Rating, review.Comment, review.DatePosted, review.ReviewID)
	if err != nil {
		log.Printf("%s: %v", ErrorReviewUpdate, err)
		return nil, err
	}

	return review, nil
}

func (s *reviewServiceImpl) DeleteReview(reviewID int) error {
	const query = `DELETE FROM reviews WHERE review_id=$1`

	_, err := s.db.Exec(query, reviewID)
	if err != nil {
		log.Printf("%s: %v", ErrorReviewDelete, err)
		return err
	}

	return nil
}

func (s *reviewServiceImpl) ListReviewsBySalonID(salonID int) ([]*models.Review, error) {
	const query = `
		SELECT review_id, user_id, salon_id, rating, comment, date_posted
		FROM reviews WHERE salon_id=$1
	`

	rows, err := s.db.Query(query, salonID)
	if err != nil {
		log.Printf("%s: %v", ErrorReviewListSalon, err)
		return nil, err
	}
	defer rows.Close()

	var reviews []*models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ReviewID, &review.UserID, &review.SalonID, &review.Rating, &review.Comment, &review.DatePosted,
		); err != nil {
			log.Printf("Error scanning review row: %v", err)
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (s *reviewServiceImpl) ListReviewsByUserID(userID int) ([]*models.Review, error) {
	const query = `
		SELECT review_id, user_id, salon_id, rating, comment, date_posted
		FROM reviews WHERE user_id=$1
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		log.Printf("%s: %v", ErrorReviewListUser, err)
		return nil, err
	}
	defer rows.Close()

	var reviews []*models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ReviewID, &review.UserID, &review.SalonID, &review.Rating, &review.Comment, &review.DatePosted,
		); err != nil {
			log.Printf("Error scanning review row: %v", err)
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (s *reviewServiceImpl) ListReviewsByRating(rating int) ([]*models.Review, error) {
	const query = `
		SELECT review_id, user_id, salon_id, rating, comment, date_posted
		FROM reviews WHERE rating=$1
	`

	rows, err := s.db.Query(query, rating)
	if err != nil {
		log.Printf("%s: %v", ErrorReviewListRating, err)
		return nil, err
	}
	defer rows.Close()

	var reviews []*models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ReviewID, &review.UserID, &review.SalonID, &review.Rating, &review.Comment, &review.DatePosted,
		); err != nil {
			log.Printf("Error scanning review row: %v", err)
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

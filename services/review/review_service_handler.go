package review

import (
	"bookmysalon/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ReviewHandler represents the HTTP handler for managing user reviews.
type ReviewHandler struct {
	service ReviewService
}

// NewReviewHandler initializes and returns an instance of ReviewHandler.
func NewReviewHandler(s ReviewService) *ReviewHandler {
	return &ReviewHandler{service: s}
}

// CreateReview creates a new review.
func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	createdReview, err := h.service.CreateReview(&review)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdReview)
}

// GetReviewByID retrieves a review by its unique ID.
func (h *ReviewHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["reviewID"])
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	review, err := h.service.GetReviewByID(reviewID)
	if err != nil {
		switch err {
		case ErrReviewNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(review)
}

// UpdateReview updates the details of an existing review.
func (h *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	updatedReview, err := h.service.UpdateReview(&review)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedReview)
}

// DeleteReview deletes a review by its unique ID.
func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["reviewID"])
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteReview(reviewID); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ListReviewsBySalonID retrieves all reviews for a specific salon by salon ID.
func (h *ReviewHandler) ListReviewsBySalonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID, err := strconv.Atoi(vars["salonID"])
	if err != nil {
		http.Error(w, "Invalid salon ID", http.StatusBadRequest)
		return
	}

	reviews, err := h.service.ListReviewsBySalonID(salonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviews)
}

// ListReviewsByUserID retrieves all reviews posted by a specific user by user ID.
func (h *ReviewHandler) ListReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	reviews, err := h.service.ListReviewsByUserID(userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviews)
}

// ListReviewsByRating retrieves all reviews with a specific rating.
func (h *ReviewHandler) ListReviewsByRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rating, err := strconv.Atoi(vars["rating"])
	if err != nil {
		http.Error(w, "Invalid rating", http.StatusBadRequest)
		return
	}

	reviews, err := h.service.ListReviewsByRating(rating)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviews)
}

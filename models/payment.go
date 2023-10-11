// bookmysalon/models/payment.go

package models

// Transaction represents a payment transaction in the system.
// swagger:model
type Transaction struct {
	// The unique ID for the transaction.
	//
	// required: true
	// example: 1001
	TransactionID int `json:"transaction_id"`

	// The ID of the user who initiated the transaction.
	//
	// required: true
	// example: 1
	UserID int `json:"user_id"`

	// The amount involved in the transaction.
	//
	// required: true
	// example: 59.99
	Amount float64 `json:"amount"`

	// The date of the transaction.
	//
	// required: true
	// example: "2023-05-20"
	Date string `json:"date"`

	// The status of the transaction.
	//
	// required: true
	// example: "Completed"
	Status string `json:"status"`

	// The method of payment used.
	//
	// required: true
	// example: "Credit Card"
	PaymentMethod string `json:"payment_method"`
}

// Invoice represents an invoice linked to a transaction in the system.
// swagger:model
type Invoice struct {
	// The unique ID for the invoice.
	//
	// required: true
	// example: 2001
	InvoiceID int `json:"invoice_id"`

	// The transaction ID associated with the invoice.
	//
	// required: true
	// example: 1001
	TransactionID int `json:"transaction_id"`

	// The details or items on the invoice.
	//
	// required: true
	// example: "Haircut, Styling"
	Details string `json:"details"`

	// The date when the invoice was issued.
	//
	// required: true
	// example: "2023-05-21"
	DateIssued string `json:"date_issued"`
}

// Promotion represents a promotional offer or discount in the system.
// swagger:model
type Promotion struct {
	// The unique ID for the promotion.
	//
	// required: true
	// example: 3001
	PromotionID int `json:"promotion_id"`

	// A description of the promotion.
	//
	// required: true
	// example: "Summer Special Discount"
	Description string `json:"description"`

	// The amount of discount provided by the promotion.
	//
	// required: true
	// example: 10.00
	DiscountAmount float64 `json:"discount_amount"`

	// The start date for the promotion's validity.
	//
	// required: true
	// example: "2023-06-01"
	ValidFrom string `json:"valid_from"`

	// The end date for the promotion's validity.
	//
	// required: true
	// example: "2023-08-31"
	ValidTo string `json:"valid_to"`
}

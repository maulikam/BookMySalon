// bookmysalon/models/payment.go

package models

type Transaction struct {
	TransactionID int     `json:"transaction_id"`
	UserID        int     `json:"user_id"`
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
}

type Invoice struct {
	InvoiceID     int    `json:"invoice_id"`
	TransactionID int    `json:"transaction_id"`
	Details       string `json:"details"`
	DateIssued    string `json:"date_issued"`
}

type Promotion struct {
	PromotionID    int     `json:"promotion_id"`
	Description    string  `json:"description"`
	DiscountAmount float64 `json:"discount_amount"`
	ValidFrom      string  `json:"valid_from"`
	ValidTo        string  `json:"valid_to"`
}

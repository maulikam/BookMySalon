package swaggerdocs

// Generic error model
// swagger:response errorResponse
type errorResponse struct {
	// The error message
	// required: true
	Message string `json:"message"`
}

// A token model
// swagger:response tokenResponse
type tokenResponse struct {
	// The JWT token
	// required: true
	Token string `json:"token"`
}

// A user representation without password
// swagger:response userResponse
type userResponse struct {
	// User's unique ID
	// required: true
	ID int `json:"id"`

	// Username for the user
	// required: true
	Username string `json:"username"`
}

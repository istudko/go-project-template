package contract

import (
	"encoding/json"
	"net/http"
)

type Ticket struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CreateTicketRequest struct {
	Name string `json:"name" db:"name" validate:"required"`
}

type CreateTicketResponse struct {
	*Ticket
}

// NewCreateTicketRequest is function to extract create ticket request body from request
func NewCreateTicketRequest(r *http.Request) (*CreateTicketRequest, error) {
	var req CreateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

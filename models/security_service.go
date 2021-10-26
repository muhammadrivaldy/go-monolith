package models

type (
	RegisterAPIRequest struct {
		ID        int64
		Name      string
		Endpoint  string
		Method    string
		ServiceID int
	}
)

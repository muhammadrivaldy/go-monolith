package models

type (
	// Pagination is a object for pagination query
	Pagination struct {
		Limit  int `json:"limit" required:"min=min=1"`
		Offset int `json:"page" required:"min=min=0"`
	}

	// PaginationResponse ..
	PaginationResponse struct {
		Limit  int   `json:"limit"`
		Offset int   `json:"page"`
		Total  int64 `json:"total"`
	}
)

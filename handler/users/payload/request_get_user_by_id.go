package payload

type RequestGetUserByID struct {
	UserID int64 `json:"user_id" validate:"required,min=1"`
}

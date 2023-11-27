package payload

type RequestGetAccessApi struct {
	UserType int `json:"user_type" validate:"required,min=1"`
}

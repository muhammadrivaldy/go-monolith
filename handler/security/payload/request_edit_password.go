package payload

type RequestEditPassword struct {
	UserID   int    `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=12"`
}

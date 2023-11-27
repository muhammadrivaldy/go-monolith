package payload

type RequestGetUserById struct {
	UserId int64 `json:"user_id" validate:"required,min=1"`
}

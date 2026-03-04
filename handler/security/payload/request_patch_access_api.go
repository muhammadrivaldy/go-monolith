package payload

type RequestPatchAccessApi struct {
	UserType int      `json:"user_type" validate:"required,min=1"`
	ApiID    []string `json:"api_id"`
}

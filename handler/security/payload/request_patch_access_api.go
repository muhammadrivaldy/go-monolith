package payload

type RequestPatchAccessApi struct {
	UserType int   `json:"user_type" validate:"required,min=1"`
	ApiId    []int `json:"api_id"`
}

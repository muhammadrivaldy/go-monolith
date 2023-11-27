package payload

type RequestVersionSupport struct {
	Version string `json:"version" validate:"required"`
}

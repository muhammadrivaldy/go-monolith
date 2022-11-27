package payload

type (
	RegisterApiRequest struct {
		ID        int64
		Name      string
		Endpoint  string
		Method    string
		ServiceID int
	}
)

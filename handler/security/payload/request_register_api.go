package payload

type RequestRegisterApi struct {
	ID        int64
	Name      string
	Endpoint  string
	Method    string
	ServiceID int64
}

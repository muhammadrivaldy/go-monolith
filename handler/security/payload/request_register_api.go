package payload

type RequestRegisterApi struct {
	ID        string
	Name      string
	Endpoint  string
	Method    string
	ServiceID int64
}

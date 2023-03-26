package payload

type RegisterApiRequest struct {
	Id        int64
	Name      string
	Endpoint  string
	Method    string
	ServiceId int
}

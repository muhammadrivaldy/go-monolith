package payload

type ResponseGetServices struct {
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"service_name"`
}

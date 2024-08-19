package payload

type ResponseGetApisServiceID struct {
	ApiID    string `json:"api_id"`
	ApiName  string `json:"api_name"`
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

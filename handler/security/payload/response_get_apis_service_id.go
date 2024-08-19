package payload

type ResponseGetApisServiceId struct {
	ApiId    string `json:"api_id"`
	ApiName  string `json:"api_name"`
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

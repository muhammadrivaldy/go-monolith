package payload

type ResponseGetApisServiceId struct {
	ApiId    int    `json:"api_id"`
	ApiName  string `json:"api_name"`
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

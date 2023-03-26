package payload

type ResponseLogin struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

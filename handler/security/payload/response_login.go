package payload

type ResponseLogin struct {
	UserId       int64  `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

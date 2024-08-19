package payload

type ResponseGetUserByID struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Status       int    `json:"status"`
	StatusName   string `json:"status_name"`
	UserType     int    `json:"user_type"`
	UserTypeName string `json:"user_type_name"`
}

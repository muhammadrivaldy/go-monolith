package config

var Config Configuration

// Configuration is a object configuration
type Configuration struct {
	Port     int    `json:"port" env:"PORT"`
	JWTKey   string `json:"jwt-key" env:"JWT_KEY"`
	Database struct {
		User          string `json:"user" env:"DB_USER"`
		Password      string `json:"password" env:"DB_PASSWORD"`
		Parameters    string `json:"parameters" env:"DB_PARAMETERS"`
		Address       string `json:"address" env:"DB_ADDRESS"`
		Database      string `json:"database" env:"DB_SCHEMA"`
		MigrationPath string `json:"migration_path" env:"DB_MIGRATION_PATH"`
	} `json:"database"`
	Redis struct {
		Address  string `json:"address" env:"REDIS_ADDRESS"`
		Password string `json:"password" env:"REDIS_PASSWORD"`
	} `json:"redis"`
	ThirdParty struct {
		Telegram struct {
			Token  string `json:"token" env:"TELEGRAM_TOKEN"`
			ChatId int64  `json:"chat_id" env:"TELEGRAM_CHAT_ID"`
		} `json:"telegram"`
	} `json:"third_party"`
	Emails struct {
		Host    string `json:"host"`
		NoReply struct {
			Email    string `json:"email"`
			Password string `json:"password" env:"NO_REPLY_PASSWORD"`
			Port     int    `json:"port"`
		} `json:"no_reply"`
	} `json:"emails"`
	Uptrace struct {
		DSN         string `json:"dsn"`
		ServiceName string `json:"service_name"`
		Environment string `json:"environment"`
	} `json:"uptrace"`
	SwaggerAPIKey string `json:"swagger_api_key" env:"SWAGGER_API_KEY"`
	BaseURL       string `json:"base_url" env:"BASE_URL"`
}

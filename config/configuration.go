package config

// Configuration is a object configuration
type Configuration struct {
	Port     int    `json:"port" env:"PORT"`
	JWTKey   string `json:"jwt-key" env:"JWT_KEY"`
	Database struct {
		User       string `json:"user" env:"DB_USER"`
		Password   string `json:"password" env:"DB_PASSWORD"`
		Parameters string `json:"parameters" env:"DB_PARAMETERS"`
		Schema     struct {
			Security struct {
				Address       string `json:"address" env:"DB_ADDRESS_SECURITY"`
				Database      string `json:"database" env:"DB_SCHEMA_SECURITY"`
				MigrationPath string `json:"migration_path" env:"DB_MIGRATION_PATH_SECURITY"`
			} `json:"security"`
			Users struct {
				Address       string `json:"address" env:"DB_ADDRESS_USERS"`
				Database      string `json:"database" env:"DB_SCHEMA_USERS"`
				MigrationPath string `json:"migration_path" env:"DB_MIGRATION_PATH_USERS"`
			} `json:"users"`
		} `json:"schema"`
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
}

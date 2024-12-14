package config

import (
	"os"
)

type Config struct {
	CredentialsFilePath string
	DSN                 string
}

func LoadConfig() *Config {
	return &Config{
		// docker-compose.yml や .env で CREDENTIALS_FILE_PATH, DB_DSN を設定
		CredentialsFilePath: os.Getenv("CREDENTIALS_FILE_PATH"),
		DSN:                 os.Getenv("DB_DSN"),
	}
}

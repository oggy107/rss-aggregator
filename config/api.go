package config

import "github.com/oggy107/rss-aggregator/internal/database"

const CONFIG_CTX = "config"
const USER_CTX = "user"

type ApiConfig struct {
	DB *database.Queries
}

func Init() *ApiConfig {
	conn := database.Connect()

	return &ApiConfig{
		DB: database.New(conn),
	}
}

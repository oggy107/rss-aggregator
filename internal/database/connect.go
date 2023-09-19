package database

import (
	"database/sql"
	"os"

	"github.com/oggy107/rss-aggregator/utils"
)

func Connect() (conn *sql.DB) {
	CONNECTION_STRING := os.Getenv("POSTGRES_CONNECTION")

	if CONNECTION_STRING == "" {
		utils.LogFatal("Database connection string is not set")
	}

	conn, err := sql.Open("postgres", CONNECTION_STRING)

	if err != nil {
		utils.LogFatal(err.Error())
	}

	return conn
}

package config

import "os"

type DbConfig struct {
	ConnString string
}

var Settings DbConfig

func init() {
	LoadEnvVariables()

	dsn := os.Getenv("DB_URL")

	Settings = DbConfig{
		ConnString: dsn,
	}
}
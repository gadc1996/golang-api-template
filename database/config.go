package database

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	User		string
	Password	string
	Database	string
	Host		string
	Protocol	string
	Port		string
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		Protocol: os.Getenv("DB_PROTOCOL"),
		Port: os.Getenv("DB_PORT"),
	}
}

func (config DatabaseConfig) FormatDSN() string {
	return fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s", 
		config.User, 
		config.Password, 
		config.Protocol, 
		config.Host, 
		config.Port, 
		config.Database,
	)
}
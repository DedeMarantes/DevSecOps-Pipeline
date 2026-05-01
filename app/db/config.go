package db

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func LoadConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	if cfg.Host == "" {
		log.Println("Missing Envinroment Variables")
	}
	return cfg
}

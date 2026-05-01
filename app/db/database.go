package db

import (
	"app/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg Config) {
	var err error
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port,
	)
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to Postgres:", err)
	}
	err = DB.AutoMigrate(&model.Anime{})
	if err != nil {
		log.Println("Failed to migrate database:", err)
	}

}

func GetDB() *gorm.DB {
	return DB
}

package db

import (
	"fmt"

	"github.com/sourav014/go-moderation-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func NewDatabase(config config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata", config.DB.Host, config.DB.Username, config.DB.Password, config.DB.DBName, config.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error while setting up database %w", err)
	}

	return &Database{
		Db: db,
	}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.Db
}

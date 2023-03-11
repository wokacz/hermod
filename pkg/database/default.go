package database

import (
	"fmt"
	"github.com/wokacz/hermod/pkg/env"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func New() (db *gorm.DB, err error) {
	config := &Config{
		Host:     env.Get("DATABASE_HOST"),
		Port:     env.Get("DATABASE_PORT"),
		Password: env.Get("DATABASE_PASSWORD"),
		User:     env.Get("DATABASE_USER"),
		SSLMode:  env.Get("DATABASE_SSLMODE"),
		DBName:   env.Get("DATABASE_DB_NAME"),
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Millisecond,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	return db, err
}

func Init() (err error) {
	DB, err = New()
	return err
}

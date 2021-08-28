package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	TimeZone string
}

func buildDBConfig() *DBConfig {

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Missing port")
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.TimeZone,
	)
}

func InitializeDB() {

	var err error

	DB, err = gorm.Open(postgres.Open(dbURL(buildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("statuse: ", err)
	}
}

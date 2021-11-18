package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBType = *gorm.DB

func New(migrate bool) (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbHost == "" {
		log.Fatal("Database URL must be provided")
	}
	if dbUser == "" {
		log.Fatal("Database User must be provided")
	}
	if dbPass == "" {
		log.Fatal("Database Password must be provided")
	}
	if sslMode == "" {
		sslMode = "disable"
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, sslMode)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if migrate {
		db.AutoMigrate(
			&Organization{},
			&Asset{},
			&Member{},
			&User{},
			&Session{},
			&Wallet{},
			&EventCapability{},
		)
	}

	return db, err
}

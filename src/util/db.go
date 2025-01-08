package util

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable for the DB connection

// ConnectDB initializes the global DB connection with connection pooling
func ConnectDB(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return err
	}

	// Get the underlying sql.DB instance to configure connection pooling
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Could not get sql.DB from Gorm DB: %v", err)
		return err
	}

	// Configure the connection pool
	sqlDB.SetMaxOpenConns(50)                 // Set the maximum number of open connections
	sqlDB.SetMaxIdleConns(25)                 // Set the maximum number of idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Set the maximum connection lifetime

	return nil
}

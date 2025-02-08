package postgresql

import (
	"codebase-go-echo/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Reduce logging noise in production
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")

	AutoMigrate()
}

// AutoMigrate ensures database tables are up-to-date
func AutoMigrate() {
	err := DB.AutoMigrate(&models.User{}) // Add more models as needed
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.Close()
	log.Println("Database connection closed")
}

// PingDB checks the database connection health
func PingDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// WithTransaction executes a function within a database transaction
func WithTransaction(txFunc func(tx *gorm.DB) error) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := txFunc(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// SetMaxConnections configures the database connection pool
func SetMaxConnections(maxOpen, maxIdle int, maxLifetime time.Duration) {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(maxOpen)        // Max open connections
	sqlDB.SetMaxIdleConns(maxIdle)        // Max idle connections
	sqlDB.SetConnMaxLifetime(maxLifetime) // Lifetime of connections
}

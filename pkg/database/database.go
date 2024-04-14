package database

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config represents the configuration for PostgreSQL connection
type Config struct {
	PsgInfo string
}

// NewDB creates a new PostgreSQL database connection
func NewDB(sqlDB *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := gormDB.Error; err != nil {
		return nil, err
	}

	return gormDB, nil
}

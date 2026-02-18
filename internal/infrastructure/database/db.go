package database

import (
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func GetSQLDB(db *gorm.DB) (*sql.DB, error) {
	return db.DB()
}

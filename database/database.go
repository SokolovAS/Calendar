package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	connection, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		return connection, err
	}

	return connection, err
}

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	connection, _ := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	return connection
}

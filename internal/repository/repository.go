package repository

import (
	"Calendar/database"
	"Calendar/entity"
	"errors"
	"gorm.io/gorm"
)

type SqliteRepo interface {
	Create(*entity.User) error
	GetEmail(email string) (entity.User, error)
}

type sqliteRepo struct {
	conn *gorm.DB
}

func NewSqliteRepo() SqliteRepo {
	return &sqliteRepo{
		conn: database.NewGormDB(),
	}
}

func (s *sqliteRepo) GetEmail(email string) (entity.User, error) {
	var user entity.User
	result := s.conn.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return user, errors.New(`"error":"Error fetching data"`)
	}
	return user, nil
}

func (s *sqliteRepo) Create(m *entity.User) error {
	result := s.conn.Create(&m)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

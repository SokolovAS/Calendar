package repository

import (
	"Calendar/database"
	"Calendar/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type SqliteRepo interface {
	Create(*entity.User)
	GetEmail(email string) (entity.User, error)
}

type sqliteRepo struct {
	gormConnection gormConnection
}

func (s *sqliteRepo) Create(user *entity.User) {
	s.gormConnection.Create(&user)
}

func (s *sqliteRepo) GetEmail(email string) (entity.User, error) {
	var user entity.User
	result := s.gormConnection.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return user, errors.New(`"error":"Error fetching data"`)
	}
	return user, nil
}

type gormConnection interface {
	Where(query interface{}, args ...interface{}) (tx gormScanner)
	Create(value interface{}) (tx *gorm.DB)
}

type gormScanner interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

func NewSqliteRepo() *sqliteRepo {
	connection, err := database.NewGormDB()
	if err != nil {
		log.Fatal("error db connection")
	}

	gc := &gC{conn: connection}

	return &sqliteRepo{
		gormConnection: gc,
	}
}

type gS struct {
	conn gormScanner
}

func (s *gS) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return s.conn.First(dest)
}

type gC struct {
	conn *gorm.DB
}

func (s *gC) Where(query interface{}, args ...interface{}) (tx gormScanner) {
	return s.conn.Where(query, args)
}

func (s *gC) Create(query interface{}) (tx *gorm.DB) {
	s.conn.Create(query)
	return
}

package calendar

import (
	"Calendar/entity"
	database "Calendar/initdb.d"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUserRecord(m entity.User) error
	HashPassword(m *entity.User) error
	CheckPassword(providedPassword string, userPassword string) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUserRecord(m entity.User) error {
	result := database.GlobalDB.Create(&m)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *userService) HashPassword(m *entity.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	if err != nil {
		return err
	}

	m.Password = string(bytes)

	return nil
}

func (s *userService) CheckPassword(providedPassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

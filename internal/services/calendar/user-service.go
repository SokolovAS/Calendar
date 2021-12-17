package calendar

import (
	"Calendar/entity"
	"Calendar/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUserRecord(m entity.User) error
	HashPassword(m *entity.User) error
	CheckPassword(providedPassword string, userPassword string) error
	GetEmail(email string) (entity.User, error)
}

type userService struct {
	repo repository.SqliteRepo
}

func NewUserService() UserService {
	return &userService{
		repo: repository.NewSqliteRepo(),
	}
}

func (s *userService) CreateUserRecord(m entity.User) error {
	s.repo.Create(&m)
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

func (s *userService) GetEmail(email string) (entity.User, error) {
	user, err := s.repo.GetEmail(email)
	if err == gorm.ErrRecordNotFound {
		return user, errors.New(`"error":"Error fetching data"`)
	}
	return user, nil
}

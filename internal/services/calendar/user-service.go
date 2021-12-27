package calendar

import (
	"Calendar/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo SqliteRepo
}

type SqliteRepo interface {
	Create(*entity.User)
	GetEmail(email string) (entity.User, error)
}

func NewUserService(r SqliteRepo) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUserRecord(m entity.User) error {
	s.repo.Create(&m)
	return nil
}

func (s *UserService) HashPassword(m *entity.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	if err != nil {
		return err
	}

	m.Password = string(bytes)

	return nil
}

func (s *UserService) CheckPassword(providedPassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return ServiceErr{
			Code:    401,
			Message: "incorrect password",
		}
	}
	return nil
}

func (s *UserService) GetEmail(email string) (entity.User, error) {
	user, err := s.repo.GetEmail(email)
	if err == gorm.ErrRecordNotFound {
		e := &ServiceErr{422, "Error fetching data"}
		return user, e
	}
	return user, nil
}

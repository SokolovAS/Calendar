package calendar

import (
	"Calendar/entity"
	"testing"
)

const testPassword = "password"
const testHash = "$2a$14$jawcknhJ41RvZJr60EB5d.2i2atYuttK5VQXMV1vqnegq3D92gzfy"
const testEmail = "email@twst.com"

type sqliteRepoMock struct{}

func (*sqliteRepoMock) Create(*entity.User) {
	return
}

func (*sqliteRepoMock) GetEmail(email string) (entity.User, error) {
	return entity.User{Email: email}, nil
}

func NewUserServiceMock() UserService {
	return &userService{
		repo: &sqliteRepoMock{},
	}
}

func TestCreateUserRecord(t *testing.T) {
	uS := NewUserServiceMock()
	m := entity.User{}
	err := uS.CreateUserRecord(m)
	if err != nil {
		t.Errorf("got %s", err)
	}
}

func TestHashPassword(t *testing.T) {
	uS := NewUserServiceMock()
	m := entity.User{}
	m.Password = testPassword
	err := uS.HashPassword(&m)
	if err != nil {
		t.Errorf("got %s", err)
	}
}

func TestCheckPassword(t *testing.T) {
	uS := NewUserServiceMock()
	m := entity.User{}
	m.Password = testPassword
	err := uS.CheckPassword(testPassword, testHash)
	if err != nil {
		t.Errorf("got %s", err)
	}
}

func TestGetEmail(t *testing.T) {
	uS := NewUserServiceMock()
	user, err := uS.GetEmail(testEmail)
	if err != nil {
		t.Errorf("got %s", err)
	}
	if user.Email != testEmail {
		if err != nil {
			t.Errorf("got %s want %s", user.Email, testEmail)
		}
	}
}

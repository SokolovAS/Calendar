package repository

import (
	"Calendar/entity"
	"gorm.io/gorm"
	"testing"
)

var emailTest = "test@gmail.com"

//type sqliteRepoMock struct {
//	conn gormConnection
//}
//
//func (*sqliteRepoMock) Create(*entity.User) {
//	return
//}
//
//func (*sqliteRepoMock) GetEmail(email string) (entity.User, error) {
//	user := entity.User{Email: email}
//	return user, nil
//}

type gCMock struct {
}

type gSMock struct {
}

func (*gSMock) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {

	return &gorm.DB{}
}

func (g *gCMock) Where(query interface{}, args ...interface{}) (tx gormScanner) {

	return &gSMock{}
}

func (g *gCMock) Create(value interface{}) (tx *gorm.DB) {

	return
}

func NewSqliteRepoMock() SqliteRepo {
	//connection, err := database.NewGormDB()
	//if err != nil {
	//	log.Fatal("error db connection")
	//}

	gc := &gCMock{}

	return &sqliteRepo{
		gormConnection: gc,
	}
}

func TestGetEmail(t *testing.T) {
	repo := NewSqliteRepoMock()
	got, _ := repo.GetEmail(emailTest)
	if emailTest != got.Email {
		t.Errorf("got %s want %s", got.Email, emailTest)
	}
}

func TestCreate(t *testing.T) {
	repo := NewSqliteRepoMock()
	user := entity.User{Email: "test@gmail.com"}
	repo.Create(&user)
}

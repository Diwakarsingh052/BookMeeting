package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func NewUserService(connectionInfo string) (*UserService, error) {
	Db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		return nil, err
	}
	Db.Exec("Use " + os.Getenv("DATABASE"))
	Db.LogMode(true)

	return &UserService{
		DB: Db,
	}, nil
}
func (us *UserService) AutoMigrate() error {
	//us.DB.DropTableIfExists(&User{})
	if err := us.DB.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}
func (us *UserService) Create(user *User) error {
	err := us.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (us *UserService) GetUser(user User) ([]User, error) {
	var u []User
	err := us.DB.Where("my_email = ?", user.Email).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (us *UserService) FetchMeetings() ([]User, error) {
	us.AutoMigrate()
	var user []User
	err := us.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

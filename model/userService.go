package model

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewUserService(connectionInfo string) (*UserService, error) {
	Db, err := gorm.Open("mysql", connectionInfo)

	if err != nil {
		return nil, err
	}

	Db.Exec("Use " + os.Getenv("DATABASE"))

	//Db.LogMode(true)

	return &UserService{
		DB: Db,
	}, nil
}

func (us *UserService) AutoMigrate() error {

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

func (us *UserService) BookMeeting(user User) error {

	fetchedUser, err := us.FetchMeetings()

	if err != nil {
		return err
	}

	for i := 0; i < len(fetchedUser); i++ {

		if (user.StartTime.After(fetchedUser[i].StartTime) && user.StartTime.Before(fetchedUser[i].EndTime)) || (user.StartTime.Equal(fetchedUser[i].StartTime) || user.StartTime.Equal(fetchedUser[i].EndTime)) {
			return errors.New("time slot already booked")
		}
		if (user.EndTime.After(fetchedUser[i].StartTime) && user.EndTime.Before(fetchedUser[i].EndTime)) || (user.EndTime.Equal(fetchedUser[i].StartTime) || user.EndTime.Equal(fetchedUser[i].EndTime)) {
			return errors.New("time slot already booked")
		}

	}

	err = us.Create(&user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeleteSlot(user User) error {
	err := us.DB.Where("start_time=? And email=?", user.StartTime, user.Email).Find(&user).Error
	if err != nil {
		return err
	}

	err = us.DB.Where("start_time=? And email=?", user.StartTime, user.Email).Delete(&user).Error

	if err != nil {
		return err
	}
	return nil
}

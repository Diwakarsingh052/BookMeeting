package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Email     string
	StartTime time.Time
	EndTime   time.Time
}

type Meeting struct {
	Email     string `json:"email"`
	StartTime string `json:"start_time"`
	Date      string `json:"date"`
	Duration  int    `json:"duration"`
}

type MeetingInfo struct {
	Date   int
	Month  int
	Year   int
	Hour   int
	Minute int
}

type UserService struct {
	DB *gorm.DB
}

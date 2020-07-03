package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Meeting struct {
	StartTime string
	Date      string
	Duration  int
}

type MeetingInfo struct {
	Date   int
	Month  int
	Year   int
	Hour   int
	Minute int
}
type User struct {
	Email     string
	StartTime time.Time
	EndTime   time.Time
}

func main() {

	m := Meeting{
		StartTime: "2:50",
		Date:      "4/7/2020",
		Duration:  20,
	}

	var DateData []string
	if strings.Contains(m.Date, "-") {
		DateData = ExtractDate(m.Date, "-")
	} else if strings.Contains(m.Date, "/") {
		DateData = ExtractDate(m.Date, "/")
	} else {
		log.Fatal("Enter Date in Valid Format")
	}
	if len(DateData) < 3 {
		log.Fatal("Enter Complete Date")
	}
	dateS, monthS, yearS := DateData[0], DateData[1], DateData[2]
	year := ValidateYear(yearS)
	month := ValidateMonth(monthS)
	date := ValidateDate(dateS, month)
	var startTime []string
	if strings.Contains(m.StartTime, ":") {
		startTime = strings.Split(m.StartTime, ":")
	} else {
		startTime = append(startTime, m.StartTime)
	}

	hour, minute := ValidateStartTime(startTime, date)
	mInfoStart := MeetingInfo{
		Date:   date,
		Month:  month,
		Year:   year,
		Hour:   hour,
		Minute: minute,
	}
	mInfoEnd := CalculateMeetingEnd(m.Duration, mInfoStart)
	t1 := User{
		Email:     "dev",
		StartTime: time.Date(mInfoStart.Year, time.Month(mInfoStart.Month), mInfoStart.Date, mInfoStart.Hour, mInfoStart.Minute, 0, 0, time.Local),
		EndTime:   time.Date(mInfoEnd.Year, time.Month(mInfoEnd.Month), mInfoEnd.Date, mInfoEnd.Hour, mInfoEnd.Minute, 0, 0, time.Local),
	}
	BookMeeting(t1)

}

func BookMeeting(user User) error {
	t1 := time.Date(2020, time.July, 15, 10, 6, 0, 0, time.Local)
	t2 := time.Date(2020, time.July, 15, 10, 20, 0, 0, time.Local)
	if (user.StartTime.After(t1) && user.StartTime.Before(t2)) || (user.StartTime.Equal(t1) || user.StartTime.Equal(t2)) {
		fmt.Println("Overlap")
		return errors.New("overlap")
	}
	if (user.EndTime.After(t1) && user.EndTime.Before(t2)) || (user.EndTime.Equal(t1) || user.EndTime.Equal(t2)) {
		fmt.Println("Overlap")
		return errors.New("overlap")
	}
	fmt.Println("Success")
	return nil

}

func ExtractDate(data, pattern string) []string {
	return strings.Split(data, pattern)
}

func ValidateYear(yearS string) int {
	year, err := strconv.Atoi(yearS)
	if err != nil {
		log.Fatal("Enter Valid Date")
	}
	if year < time.Now().Year() {
		log.Fatal("Date is not valid ")
	}
	return year

}
func ValidateMonth(monthS string) int {
	month, err := strconv.Atoi(monthS)
	if err != nil {
		log.Fatal("Enter Valid Date")

	}
	if time.Month(month) < time.Now().Month() {
		log.Fatal("Date is not valid")
	}
	return month
}

func ValidateDate(dateS string, month int) int {
	date, err := strconv.Atoi(dateS)
	if err != nil {
		log.Fatal("Enter Valid Date")
	}

	if date < time.Now().Day() {
		if time.Month(month) == time.Now().Month() {
			log.Fatal("Date is not valid")
		}
	}
	if date > 31 {
		log.Fatal("Date is not valid")
	}
	return date
}

func ValidateStartTime(startTime []string, date int) (int, int) {
	hour, err := strconv.Atoi(startTime[0])
	if err != nil {
		log.Fatal("Enter Valid Time")
	}

	if hour > 23 || hour < 0 {
		log.Fatal("Enter valid Time")
	}
	if (hour >= 5 && hour < 9) || hour >= 17 {
		log.Fatal("Please Select time After 9 and Before 5 only")
	}
	if hour < time.Now().Hour() && time.Now().Day() == date {
		log.Fatal("Enter valid Time")
	}
	if hour < 9 {
		hour = hour + 12
	}

	var minute int
	if len(startTime) > 1 {
		minute, err = strconv.Atoi(startTime[1])

		if err != nil || minute > 60 || minute < 0 {
			log.Fatal("Enter valid Time")
		}
	}
	return hour, minute
}
func CalculateMeetingEnd(duration int, startInfo MeetingInfo) MeetingInfo {

	startInfo.Minute = startInfo.Minute + duration
	fmt.Println(startInfo.Minute)
	if startInfo.Minute >= 60 {
		startInfo.Hour = startInfo.Hour + 1
		startInfo.Minute = startInfo.Minute - 60
	}
	fmt.Println(startInfo.Hour, startInfo.Minute)
	return startInfo
}

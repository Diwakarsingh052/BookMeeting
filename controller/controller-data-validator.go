package controller

import (
	"BookMeeting/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ValidDateError = errors.New("date is not valid")
var ValidTimeError = errors.New("enter valid time")

func ExtractDate(data, pattern string) []string {
	return strings.Split(data, pattern)
}

func ValidateYear(yearS string) (int, error) {
	year, err := strconv.Atoi(yearS)
	if err != nil {
		return 0, ValidDateError
	}
	if year < time.Now().Year() {
		return 0, ValidDateError
	}
	return year, nil

}
func ValidateMonth(monthS string, year int) (int, error) {
	month, err := strconv.Atoi(monthS)

	if err != nil {
		return 0, ValidDateError

	}

	if time.Month(month) < time.Now().Month() {
		if year == time.Now().Year() {
			return 0, ValidDateError
		}
	}
	return month, nil
}

func ValidateDate(dateS string, month int, year int) (int, error) {
	date, err := strconv.Atoi(dateS)
	if err != nil {
		return 0, ValidDateError
	}

	if date < time.Now().Day() {
		if time.Month(month) == time.Now().Month() && year == time.Now().Year() {
			return 0, ValidDateError
		}
	}

	if date > 31 {
		return 0, ValidDateError
	}
	return date, nil
}

func ValidateStartTime(startTime []string, date int, month int, year int) (int, int, error) {
	hour, err := strconv.Atoi(startTime[0])
	if err != nil {
		return 0, 0, ValidTimeError
	}

	if hour > 23 || hour < 0 {
		return 0, 0, ValidTimeError
	}

	if (hour >= 5 && hour < 9) || hour >= 17 {
		return 0, 0, errors.New("please Select time After 9 am and Before 5 pm only")
	}

	if hour < time.Now().Hour() && time.Now().Day() == date && time.Month(month) == time.Now().Month() && year == time.Now().Year() {
		return 0, 0, ValidTimeError
	}

	if hour < 9 {
		hour = hour + 12
	}

	var minute int

	if len(startTime) > 1 {
		minute, err = strconv.Atoi(startTime[1])

		if err != nil || minute > 60 || minute < 0 {
			return 0, 0, ValidTimeError
		}
	}

	return hour, minute, nil
}
func CalculateMeetingEnd(duration int, startInfo model.MeetingInfo) (model.MeetingInfo, error) {

	startInfo.Minute = startInfo.Minute + duration

	if startInfo.Minute >= 60 {
		startInfo.Hour = startInfo.Hour + 1
		startInfo.Minute = startInfo.Minute - 60
	}
	if (startInfo.Hour >= 5 && startInfo.Hour < 9) || startInfo.Hour >= 17 {
		return model.MeetingInfo{}, errors.New("you can book appointment before 5 pm only")
	}
	fmt.Println(startInfo.Hour, startInfo.Minute)
	return startInfo, nil
}

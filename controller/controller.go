package controller

import (
	"BookMeeting/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Users struct {
	us *model.UserService
}

func NewUsers(us *model.UserService) *Users {
	return &Users{
		us: us,
	}
}
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	var m model.Meeting
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(b, &m)

	var DateData []string

	if strings.Contains(m.Date, "-") {
		DateData = ExtractDate(m.Date, "-")
	} else if strings.Contains(m.Date, "/") {
		DateData = ExtractDate(m.Date, "/")
	} else {
		http.Error(w, "enter date in valid format", http.StatusBadRequest)
		return
	}

	if len(DateData) < 3 {
		http.Error(w, "enter date in valid format", http.StatusBadRequest)
		return
	}

	dateS, monthS, yearS := DateData[0], DateData[1], DateData[2]
	year, err := ValidateYear(yearS)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	month, err := ValidateMonth(monthS, year)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := ValidateDate(dateS, month, year)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var startTime []string

	if strings.Contains(m.StartTime, ":") {
		startTime = strings.Split(m.StartTime, ":")
	} else {
		startTime = append(startTime, m.StartTime)
	}

	hour, minute, err := ValidateStartTime(startTime, date, month, year)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mInfoStart := model.MeetingInfo{
		Date:   date,
		Month:  month,
		Year:   year,
		Hour:   hour,
		Minute: minute,
	}

	t1 := model.User{
		Email:     m.Email,
		StartTime: time.Date(mInfoStart.Year, time.Month(mInfoStart.Month), mInfoStart.Date, mInfoStart.Hour, mInfoStart.Minute, 0, 0, time.Local),
	}

	err = u.us.DeleteSlot(t1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("Successfully Deleted"))
}

func (u *Users) BookSLot(w http.ResponseWriter, r *http.Request) {
	var m model.Meeting
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(b, &m)

	var DateData []string

	if strings.Contains(m.Date, "-") {
		DateData = ExtractDate(m.Date, "-")
	} else if strings.Contains(m.Date, "/") {
		DateData = ExtractDate(m.Date, "/")
	} else {
		http.Error(w, "enter date in valid format", http.StatusBadRequest)
		return
	}

	if len(DateData) < 3 {
		http.Error(w, "enter date in valid format", http.StatusBadRequest)
		return
	}

	dateS, monthS, yearS := DateData[0], DateData[1], DateData[2]
	year, err := ValidateYear(yearS)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	month, err := ValidateMonth(monthS, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := ValidateDate(dateS, month, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var startTime []string
	if strings.Contains(m.StartTime, ":") {
		startTime = strings.Split(m.StartTime, ":")
	} else {
		startTime = append(startTime, m.StartTime)
	}

	hour, minute, err := ValidateStartTime(startTime, date, month, year)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mInfoStart := model.MeetingInfo{
		Date:   date,
		Month:  month,
		Year:   year,
		Hour:   hour,
		Minute: minute,
	}

	mInfoEnd, err := CalculateMeetingEnd(m.Duration, mInfoStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t1 := model.User{
		Email:     m.Email,
		StartTime: time.Date(mInfoStart.Year, time.Month(mInfoStart.Month), mInfoStart.Date, mInfoStart.Hour, mInfoStart.Minute, 0, 0, time.Local),
		EndTime:   time.Date(mInfoEnd.Year, time.Month(mInfoEnd.Month), mInfoEnd.Date, mInfoEnd.Hour, mInfoEnd.Minute, 0, 0, time.Local),
	}

	err = u.us.BookMeeting(t1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Success"))

}

func (u *Users) AllBookings(w http.ResponseWriter, r *http.Request) {
	var user []model.User

	err := u.us.DB.Find(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.MarshalIndent(&user, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write(j)
}

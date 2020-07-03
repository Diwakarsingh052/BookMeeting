package main

import (
	"BookMeeting/controller"
	"BookMeeting/model"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	us, err := model.NewUserService(os.Getenv("NAME") + ":" + os.Getenv("PASSWORD") + "@/" + os.Getenv("DATABASE") + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer us.DB.Close()
	usersC := controller.NewUsers(us)
	http.HandleFunc("/book", usersC.BookSLot)
	http.HandleFunc("/delete", usersC.Delete)
	http.HandleFunc("/all", usersC.AllBookings)
	http.ListenAndServe(":8080", nil)
}

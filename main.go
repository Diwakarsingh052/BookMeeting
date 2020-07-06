package main

import (
	"BookMeeting/controller"
	"BookMeeting/model"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

package main

import (
	"fmt"
	"github.com/SSunSShine/travel/database"
	"github.com/SSunSShine/travel/conf"
	"github.com/SSunSShine/travel/model"
	"time"
)


func initAll(conf *conf.Configuration) {

	if (database.DB.HasTable(&model.Customer{})) {
		fmt.Println("db has the table Customer, so drop it.")
		database.DB.DropTable(&model.Customer{})
	}

	if (database.DB.HasTable(&model.City{})) {
		fmt.Println("db has the table City, so drop it.")
		database.DB.DropTable(&model.City{})
	}

	if (database.DB.HasTable(&model.Car{})) {
		fmt.Println("db has the table Car, so drop it.")
		database.DB.DropTable(&model.Car{})
	}

	if (database.DB.HasTable(&model.Hotel{})) {
		fmt.Println("db has the table Hotel, so drop it.")
		database.DB.DropTable(&model.Hotel{})
	}

	if (database.DB.HasTable(&model.Flight{})) {
		fmt.Println("db has the table Flight, so drop it.")
		database.DB.DropTable(&model.Flight{})
	}

	if (database.DB.HasTable(&model.Reservation{})) {
		fmt.Println("db has the table Reservation, so drop it.")
		database.DB.DropTable(&model.Reservation{})
	}

	database.DB.AutoMigrate(&model.City{})
	database.DB.AutoMigrate(&model.Customer{})
	database.DB.AutoMigrate(&model.Car{})
	database.DB.AutoMigrate(&model.Hotel{})
	database.DB.AutoMigrate(&model.Flight{})
	database.DB.AutoMigrate(&model.Reservation{})

	city0 := model.City{CityName: "苏州"}
	city0.Create()
	city1 := model.City{CityName: "福州"}
	city1.Create()
	city2 := model.City{CityName: "桂林"}
	city2.Create()

	customer0 := model.Customer{CustName: "admin", Password: "123456", Type: 0}
	customer0.Create()

	car0 := model.Car{CarNum: "C00001", Price: 10, CityName: "桂林"}
	car0.Create()

	h0 := model.Hotel{HotelName: "FirstHotel", Price: 100, NumRooms: 100, CityName: "苏州"}
	h0.Create()

	f0 := model.Flight{FlightNum: "F00001", Price: 300, NumSeats: 50, FromCity: "苏州", ArivCity: "福州"}
	f0.Create()

	r0 := model.Reservation{ResvKey: "C00001", CustName: "admin", Type: 1, ResDate: time.Now()}
	r0.Create()
	r1 := model.Reservation{ResvKey: "F00001", CustName: "admin", Type: 2, ResDate: time.Now()}
	r1.Create()

	fmt.Println("restarted success !")
}

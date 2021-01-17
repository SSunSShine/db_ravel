package model

import (
	"github.com/SSunSShine/travel/database"
	"github.com/jinzhu/gorm"
	"log"
)

type Flight struct {
	FlightNum string `json:"flight_num" gorm:"type:varchar(50);primary_key;not null"`
	Price     int    `json:"price" gorm:"type:int;DEFAULT:0"`
	NumSeats  int    `json:"num_seats" gorm:"type:int;DEFAULT:0"`
	FromCity  string `json:"from_city" gorm:"type:varchar(50);not null"`
	FCity     City   `json:"f_city" gorm:"ForeignKey:FromCity"`
	ArivCity  string `json:"ariv_city" gorm:"type:varchar(50);not null"`
	ACity     City   `json:"a_city" gorm:"ForeignKey:ArivCity"`
}

func (f *Flight) Get() (flight Flight, err error) {

	if err = database.DB.Where(&f).First(&flight).Error; err != nil {
		log.Print(err)
	}

	return
}

func (f *Flight) Create() (err error) {

	if err = database.DB.Create(&f).Error; err != nil {
		log.Print(err)
	}

	return
}

func (f *Flight) Update() (err error) {

	if err = database.DB.Model(&f).Updates(f).Error; err != nil {
		log.Print(err)
	}

	return
}

func (f *Flight) Delete() (err error) {

	if err = database.DB.Unscoped().Delete(&f).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetList func
func (f *Flight) GetList() (flights []Flight, err error) {

	if err = database.DB.Where(&f).Find(&flights).Error; err != nil {
		log.Print(err)
	}
	return
}

func (f *Flight) AfterDelete(db *gorm.DB) (err error) {
	var resv Reservation
	resv.ResvKey = f.FlightNum

	if err = db.Unscoped().Delete(&resv).Error; err != nil {
		log.Print(err)
	}
	return
}

func (f *Flight) AfterCreate() (err error) {
	var city City
	city.CityName = f.FromCity

	if _, err := city.Get(); err != nil {
		if err = city.Create(); err != nil {
			log.Print(err)
		}
	}

	city.CityName = f.ArivCity
	if _, err := city.Get(); err != nil {
		if err = city.Create(); err != nil {
			log.Print(err)
		}
	}
	return
}
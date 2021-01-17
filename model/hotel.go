package model

import (
	"github.com/SSunSShine/travel/database"
	"github.com/jinzhu/gorm"
	"log"
)

type Hotel struct {
	HotelName string `json:"hotel_name" gorm:"type:varchar(50);primary_key;not null"`
	Price     int    `json:"price" gorm:"type:int;DEFAULT:0"`
	NumRooms  int    `json:"num_rooms" gorm:"type:int;DEFAULT:0"`
	CityName  string `json:"city_name" gorm:"type:varchar(50);not null"`
	City      City   `json:"city" gorm:"ForeignKey:CityName"`
}

func (h *Hotel) Get() (hotel Hotel, err error) {

	if err = database.DB.Where(&h).First(&hotel).Error; err != nil {
		log.Print(err)
	}

	return
}

func (h *Hotel) Create() (err error) {

	if err = database.DB.Create(&h).Error; err != nil {
		log.Print(err)
	}

	return
}

func (h *Hotel) Update() (err error) {

	if err = database.DB.Model(&h).Updates(h).Error; err != nil {
		log.Print(err)
	}

	return
}

func (h *Hotel) Delete() (err error) {

	if err = database.DB.Unscoped().Delete(&h).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetList func
func (h *Hotel) GetList() (hotels []Hotel, err error) {

	if err = database.DB.Where(&h).Find(&hotels).Error; err != nil {
		log.Print(err)
	}
	return
}

func (h *Hotel) AfterDelete(db *gorm.DB) (err error) {
	var resv Reservation
	resv.ResvKey = h.HotelName

	if err = db.Unscoped().Delete(&resv).Error; err != nil {
		log.Print(err)
	}
	return
}

func (h *Hotel) AfterCreate() (err error) {
	var city City
	city.CityName = h.CityName

	if _, err := city.Get(); err != nil {
		if err = city.Create(); err != nil {
			log.Print(err)
		}
	}
	return
}
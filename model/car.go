package model

import (
	"github.com/SSunSShine/travel/database"
	"github.com/jinzhu/gorm"
	"log"
)

type Car struct {
	CarNum   string `json:"car_num" gorm:"type:varchar(50);primary_key;not null"`
	Price    int    `json:"price" gorm:"type:varchar(50);DEFAULT:0;not null"`
	CityName string `json:"city_name" gorm:"type:varchar(50);not null"`
	City     City   `json:"city" gorm:"ForeignKey:CityName"`
}

func (c *Car) Get() (car Car, err error) {

	if err = database.DB.Where(&c).First(&car).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *Car) Create() (err error) {

	if err = database.DB.Create(&c).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *Car) Update() (err error) {

	if err = database.DB.Model(&c).Updates(c).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *Car) Delete() (err error) {

	if err = database.DB.Unscoped().Delete(&c).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetList func
func (c *Car) GetList() (cars []Car, err error) {

	if err = database.DB.Where(&c).Find(&cars).Error; err != nil {
		log.Print(err)
	}
	return
}

func (c *Car) AfterDelete(db *gorm.DB) (err error) {
	var resv Reservation
	resv.ResvKey = c.CarNum

	if err = db.Unscoped().Delete(&resv).Error; err != nil {
		log.Print(err)
	}
	return
}

func (c *Car) AfterCreate() (err error) {
	var city City
	city.CityName = c.CityName
	if _, err := city.Get(); err != nil {
		if err = city.Create(); err != nil {
			log.Print(err)
		}
	}
	return
}
package model

import (
	"github.com/SSunSShine/travel/database"
	"log"
)

type City struct {
	CityName string `json:"city_name" gorm:"type:varchar(20);primary_key;not null"`
}

func (c *City) Get() (city City, err error) {

	if err = database.DB.Where(&c).First(&city).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *City) Create() (err error) {

	if err = database.DB.Create(&c).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *City) Update() (err error) {

	if err = database.DB.Model(&c).Updates(c).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *City) Delete() (err error) {

	if err = database.DB.Unscoped().Delete(&c).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetList func
func (c *City) GetList() (cities []City, err error) {

	if err = database.DB.Where(&c).Find(&cities).Error; err != nil {
		log.Print(err)
	}
	return
}
package model

import (
	"github.com/SSunSShine/travel/database"
	"log"
	"time"
)

type Reservation struct {
	ResvKey  string    `json:"resv_key" gorm:"type:varchar(50);primary_key;not null"`
	CustName string    `json:"cust_name" gorm:"type:varchar(50);primary_key;not null"`
	Type     int       `json:"type" gorm:"type:int;not null"` // 1 为 Car, 2 为 Flight, 3 为 Hotel
	ResDate  time.Time `json:"res_date" gorm:"type:DATE;primary_key;not null"`
}

func (r *Reservation) Get() (reservation Reservation, err error) {

	if err = database.DB.Where(&r).First(&reservation).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reservation) Create() (err error) {

	if err = database.DB.Create(&r).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reservation) Update() (err error) {

	if err = database.DB.Model(&r).Updates(r).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reservation) Delete() (err error) {

	if err = database.DB.Unscoped().Delete(&r).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetList func
func (r *Reservation) GetList() (reservations []Reservation, err error) {

	if err = database.DB.Where(&r).Order("res_date asc").Find(&reservations).Error; err != nil {
		log.Print(err)
	}
	return
}
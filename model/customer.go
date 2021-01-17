package model

import (
	"github.com/SSunSShine/travel/database"
	"log"
)

type Customer struct {
	CustName string `json:"cust_name" gorm:"type:varchar(50);primary_key;not null"`
	Password string `json:"password" gorm:"type:varchar(50);not null"`
	Type     int    `json:"type" gorm:"type:int;DEFAULT:0"`
}

func (c *Customer) Get() (customer Customer, err error) {
	if err = database.DB.Where(&c).First(&customer).Error; err != nil {
		log.Print(err)
	}
	return
}
func (c *Customer) Create() (err error) {

	if err = database.DB.Create(&c).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *Customer) Update() (err error) {

	if err = database.DB.Model(&c).Updates(c).Error; err != nil {
		log.Print(err)
	}

	return
}

func (c *Customer) Delete() (err error) {

	if err = database.DB.Unscoped().Delete(&c).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetList func
func (c *Customer) GetList() (customers []Customer, err error) {

	if err = database.DB.Where(&c).Find(&customers).Error; err != nil {
		log.Print(err)
	}
	return
}
package controller

import (
	"github.com/SSunSShine/travel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCar
func GetCar(c *gin.Context) {

	var car model.Car

	car.CarNum = c.Param("carNum")

	getCar, err := car.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": getCar,
	})

}

// UpdateCar
func UpdateCar(c *gin.Context)  {

	var car model.Car

	car.CarNum = c.Param("carNum")

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}


	if err := car.Update(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "success",
		})
	}
}

// DeleteCar
func DeleteCar(c *gin.Context)  {

	var car model.Car

	car.CarNum = c.Param("CarNum")

	getCar, err := car.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if err := getCar.Delete(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "success",
		})
	}
}

// CreateCar
func CreateCar(c *gin.Context)  {

	var car model.Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if err := car.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": car,
	})
}

// GetCars
func GetCars(c *gin.Context) {

	var car model.Car

	car.CityName = c.Query("cityName")

	cars, err := car.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": cars,
	})

}
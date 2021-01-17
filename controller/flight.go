package controller

import (
	"github.com/SSunSShine/travel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetFlight
func GetFlight(c *gin.Context) {

	var f model.Flight

	f.FlightNum = c.Param("flightNum")

	flight, err := f.Get()
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
		"data": flight,
	})

}

// UpdateFlight
func UpdateFlight(c *gin.Context)  {

	var f model.Flight

	f.FlightNum = c.Param("flightNum")

	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}


	if err := f.Update(); err != nil {
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

// DeleteFlight
func DeleteFlight(c *gin.Context)  {

	var f model.Flight

	f.FlightNum = c.Param("flightNum")

	flight, err := f.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if err := flight.Delete(); err != nil {
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

// CreateFlight
func CreateFlight(c *gin.Context)  {

	var f model.Flight

	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if err := f.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": f,
	})
}

// GetFlights
func GetFlights(c *gin.Context) {

	var f model.Flight

	f.ArivCity = c.Query("toCity")
	f.FromCity = c.Query("fromCity")

	flights, err := f.GetList()
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
		"data": flights,
	})

}
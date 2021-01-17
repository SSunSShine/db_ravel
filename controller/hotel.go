package controller

import (
	"github.com/SSunSShine/travel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHotel
func GetHotel(c *gin.Context) {

	var h model.Hotel

	h.HotelName = c.Param("hotelName")

	hotel, err := h.Get()
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
		"data": hotel,
	})

}

// UpdateHotel
func UpdateHotel(c *gin.Context)  {

	var h model.Hotel

	h.HotelName = c.Param("hotelName")

	if err := c.ShouldBindJSON(&h); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}


	if err := h.Update(); err != nil {
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

// DeleteHotel
func DeleteHotel(c *gin.Context)  {

	var h model.Hotel

	h.HotelName = c.Param("hotelName")

	hotel, err := h.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if err := hotel.Delete(); err != nil {
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

// CreateHotel
func CreateHotel(c *gin.Context)  {

	var h model.Hotel

	if err := c.ShouldBindJSON(&h); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if err := h.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": h,
	})
}

// GetHotels
func GetHotels(c *gin.Context) {

	var h model.Hotel

	h.CityName = c.Query("cityName")

	hotels, err := h.GetList()
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
		"data": hotels,
	})

}
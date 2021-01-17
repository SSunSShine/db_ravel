package controller

import (
	"github.com/SSunSShine/travel/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ReservationVO struct {
	ResvKey   string `json:"resv_key"`
	CustName  string `json:"cust_name"`
	Type      int    `json:"type"` // 1 为 Car, 2 为 Flight, 3 为 Hotel
	Price     int    `json:"price"`
	From_city string `json:"from_city"`
	Ariv_city string `json:"ariv_city"`
	ResDate   string `json:"res_date"`
}

// GetReservation
func GetReservation(c *gin.Context) {

	var r model.Reservation

	r.ResvKey = c.Param("resvKey")

	reservation, err := r.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    reservation,
	})

}

func UpdateReservation(c *gin.Context) {

	var r model.Reservation

	r.ResvKey = c.Param("resvKey")

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if err := r.Update(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})
	}
}

// DeleteReservation
func DeleteReservation(c *gin.Context) {

	var r model.Reservation

	r.ResvKey = c.Param("resvKey")

	reservation, err := r.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if err := reservation.Delete(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})
	}
}

// CreateReservation
func CreateReservation(c *gin.Context) {

	var r model.Reservation
	var cust model.Customer

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	cust.CustName = r.CustName
	_, err := cust.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if err := r.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": "您已预订该项目，无法重新预订!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    r,
	})
}

// GetReservations
func GetReservations(c *gin.Context) {

	var r model.Reservation

	r.CustName = c.Query("cust_name")

	resv, err := r.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    resv,
	})

}

// GetCarResv
func GetCarResv(c *gin.Context) {

	var r model.Reservation
	var car model.Car

	r.CustName = c.Query("cust_name")
	r.Type = 1

	resv, err := r.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	var res []ReservationVO
	for _, v := range resv {
		var resvVO ReservationVO
		resvVO.ResvKey = v.ResvKey
		resvVO.CustName = v.CustName
		resvVO.ResDate = v.ResDate.Format("2006-01-02")

		car.CarNum = v.ResvKey
		ca, err := car.Get()
		if err != nil {
			log.Print(err)
		}
		resvVO.From_city = ca.CityName
		resvVO.Ariv_city = ca.CityName
		resvVO.Price = ca.Price

		res = append(res, resvVO)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    res,
	})
}

// GetFlightResv
func GetFlightResv(c *gin.Context) {

	var r model.Reservation
	var f model.Flight

	r.CustName = c.Query("cust_name")
	r.Type = 2

	resv, err := r.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	var res []ReservationVO
	for _, v := range resv {
		var resvVO ReservationVO
		resvVO.ResvKey = v.ResvKey
		resvVO.CustName = v.CustName
		resvVO.ResDate = v.ResDate.Format("2006-01-02")

		f.FlightNum = v.ResvKey
		flight, err := f.Get()
		if err != nil {
			log.Print(err)
		}
		resvVO.From_city = flight.FromCity
		resvVO.Ariv_city = flight.ArivCity
		resvVO.Price = flight.Price

		res = append(res, resvVO)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    res,
	})

}

// GetHotelResv
func GetHotelResv(c *gin.Context) {

	var r model.Reservation
	var h model.Hotel

	r.CustName = c.Query("cust_name")
	r.Type = 3

	resv, err := r.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	var res []ReservationVO
	for _, v := range resv {
		var resvVO ReservationVO
		resvVO.ResvKey = v.ResvKey
		resvVO.CustName = v.CustName
		resvVO.ResDate = v.ResDate.Format("2006-01-02")

		h.HotelName = v.ResvKey
		hotel, err := h.Get()
		if err != nil {
			log.Print(err)
		}
		resvVO.From_city = hotel.CityName
		resvVO.Ariv_city = hotel.CityName
		resvVO.Price = hotel.Price

		res = append(res, resvVO)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    res,
	})

}

// GetRoute
func GetRoute(c *gin.Context) {

	var r model.Reservation
	var f model.Flight
	var car model.Car
	var h model.Hotel

	r.CustName = c.Query("cust_name")

	resv, err := r.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	var res []string
	for _, v := range resv {
		resvDate := v.ResDate.Format("2006-01-02")
		s := ""
		if v.Type == 1 {
			car.CarNum = v.ResvKey
			ca, err := h.Get()
			if err != nil {
				log.Print(err)
			}
			s += "您预订的位于 " + ca.CityName + " 的专车: " + car.CarNum
			s += " 乘坐时间为: " + resvDate
		} else if v.Type == 2 {
			f.FlightNum = v.ResvKey
			flight, err := f.Get()
			if err != nil {
				log.Print(err)
			}
			s += "您预订的从 " + flight.FromCity + " 飞往 " + flight.ArivCity + " 的航班: " + flight.FlightNum
			s += " 乘坐时间为: " + resvDate
		} else {
			h.HotelName = v.ResvKey
			hotel, err := h.Get()
			if err != nil {
				log.Print(err)
			}
			s += "您预订的位于 " + hotel.CityName + " 的宾馆: " + hotel.HotelName
			s += " 居住时间为: " + resvDate
		}

		s += "。"
		res = append(res, s)
	}
	res = append(res, "请勿错过时间！逾期不退！谢谢合作！")

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    res,
	})

}

package route

import (
	"github.com/SSunSShine/travel/controller"
	"github.com/SSunSShine/travel/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	auth := r.Group("/api")
	auth.Use(middleware.JwtToken())
	{
		// customer
		auth.PUT("customer/:custName", controller.UpdateCustomer)
		auth.DELETE("customer/:custName", controller.DeleteCustomer)

		// car
		auth.PUT("car/:carNum", controller.UpdateCar)
		auth.DELETE("car/:carNum", controller.DeleteCar)
		auth.POST("car/addCar", controller.CreateCar)

		// flight
		auth.POST("flight/addFlight", controller.CreateFlight)
		auth.PUT("flight/:flightNum", controller.UpdateFlight)
		auth.DELETE("flight/:flightNum", controller.DeleteFlight)

		// hotel
		auth.POST("hotel/addHotel", controller.CreateHotel)
		auth.PUT("hotel/:hotelName", controller.UpdateHotel)
		auth.DELETE("hotel/:hotelName", controller.DeleteHotel)

		// reservation
		auth.POST("reservation/addReservation", controller.CreateReservation)
		auth.GET("reservation/:resvKey", controller.GetReservation)
		auth.GET("reservations/list", controller.GetReservations)
		auth.PUT("reservation/:resvKey", controller.UpdateReservation)
		auth.DELETE("reservation/:resvKey", controller.DeleteReservation)
		auth.GET("reservations/carList", controller.GetCarResv)
		auth.GET("reservations/flightList", controller.GetFlightResv)
		auth.GET("reservations/hotelList", controller.GetHotelResv)
		auth.GET("reservations/routeList", controller.GetRoute)
	}
	router := r.Group("/api")
	{
		router.POST("sign", controller.Sign)
		router.POST("login", controller.Login)

		// customer
		router.GET("customer/:custName", controller.GetCustomer)
		router.GET("customers/list", controller.GetCustomers)

		// car
		router.GET("car/:carNum", controller.GetCar)
		router.GET("cars/list", controller.GetCars)

		// flight
		router.GET("flight/:flightNum", controller.GetFlight)
		router.GET("flights/list", controller.GetFlights)

		// hotel
		router.GET("hotel/:hotelName", controller.GetHotel)
		router.GET("hotels/list", controller.GetHotels)
	}
}
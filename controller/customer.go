package controller

import (
	"github.com/SSunSShine/travel/middleware"
	"github.com/SSunSShine/travel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCustomer 获取单个用户信息
func GetCustomer(c *gin.Context) {

	var cust model.Customer

	cust.CustName = c.Param("custName")

	customer, err := cust.Get()
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
		"data": customer,
	})

}

// UpdateCustomer 更新用户信息。 注：更新Profile个人简介需要调用UpdateProfile
func UpdateCustomer(c *gin.Context)  {

	var cust model.Customer

	cust.CustName = c.Param("custName")

	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}


	if err := cust.Update(); err != nil {
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

// DeleteCustomer
func DeleteCustomer(c *gin.Context)  {

	var cust model.Customer

	cust.CustName = c.Param("custName")

	customer, err := cust.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if err := customer.Delete(); err != nil {
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

func Sign(c *gin.Context)  {

	var cust model.Customer

	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if err := cust.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "success",
		})
	}
}

func Login(c *gin.Context)  {

	var cust model.Customer

	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	customer, err := cust.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if cust.CustName != customer.CustName || cust.Password != customer.Password {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"message": "用户名或密码不正确!",
		})
		return
	}
	token, err := middleware.Gen(customer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"token": token,
		"custName": customer.CustName,
	})
}

// GetCustomers
func GetCustomers(c *gin.Context) {

	var cust model.Customer

	custs, err := cust.GetList()
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
		"data": custs,
	})

}
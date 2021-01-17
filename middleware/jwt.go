package middleware

import (
	"errors"
	"github.com/SSunSShine/travel/conf"
	"github.com/SSunSShine/travel/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	CustName  string `json:"cust_name"`
	jwt.StandardClaims
}

var JwtKey = []byte(conf.Config().JwtKey)

// Gen 生成token
func Gen(customer model.Customer) (token string, err error) {

	// token 存活时间 10h
	expireTime := time.Now().Add(48 * time.Hour)
	SetClaims := MyClaims{
		customer.CustName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "localhost",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err = reqClaim.SignedString(JwtKey)
	if err != nil {
		log.Print(err)
	}

	return
}

// Parse 解析token
func Parse(token string) (custName string, err error) {

	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("Token not valid")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				err = errors.New("Token has expired")
			} else {
				err = errors.New("Incorrect Token format")
			}
		}
		return
	}

	if info, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		custName = info.CustName
	} else {
		err = errors.New("Token not valid")
	}
	return
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		// 请求头信息得到token
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			err = errors.New("Token does not exist")
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			err = errors.New("Incorrect Token format")
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			err = errors.New("Incorrect Token format")
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		custName, err := Parse(checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("custName", custName)

		c.Next()

		log.Print("custName:", custName)
	}
}

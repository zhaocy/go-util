package go_util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//Infof("Req: %#v",c.Request.Body)
		c.Next()
	}
}

func NoResponse(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET,OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		c.Next()
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "API not exists!",
		})
	}

}

type JSONErr struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		ct := c.GetHeader("Content-Type")
		authToken := c.GetHeader("Authorization")
		if ct == "" || ct != "application/json" {
			c.JSON(http.StatusOK, &JSONErr{
				Code: 310,
				Msg:  "not support Content-Type",
			})
			c.Abort()
		} else if authToken == "" {
			c.JSON(http.StatusOK, &JSONErr{
				Code: 410,
				Msg:  "missing Authorization",
			})
			c.Abort()
		}
		c.Next()
	}
}

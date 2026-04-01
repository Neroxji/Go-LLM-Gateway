package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 全局通用跨域
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 鉴权中间件
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		
		// apikey:=c.GetHeader("Authorization")



	}
}
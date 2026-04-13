// Package middleware 跨域中间件配置
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 直接设置允许的前端域名
		c.Header("Access-Control-Allow-Origin", "https://admin.010814.xyz")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, AccessToken, X-CSRF-Token, Authorization, Token, Accept, Accept-Encoding, Origin, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")
		
		// 放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		// 处理请求
		c.Next()
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"medicine/utils"
	"net/http"
)

// 需要放行的 API 白名单
var whiteList = map[string]bool{
	"/medicine/demo/testGet":    true,
	"/medicine/demo/testPost":   true,
	"/medicine/actuator/health": true,
	"/medicine/user/login":      true,
	"/medicine/user/loginV2":    true,
}

// Jwt 鉴权中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求路径
		requestPath := c.Request.URL.Path

		// 如果在白名单中，则直接放行
		if _, exists := whiteList[requestPath]; exists {
			c.Next()
			return
		}

		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token required",
				"code":  3045})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  3046,
				"error": "无效的 token"})
			c.Abort()
			return
		}

		// 将用户名存入 context，后续处理可以使用
		c.Set("account", claims.Account)
		c.Next()
	}
}

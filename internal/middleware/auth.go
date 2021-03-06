package middleware

import (
	"github.com/gin-gonic/gin"
	"graduation/internal/config"
	"graduation/pkg/service"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "username cookie expired")
			return
		}
		session_id, err := c.Cookie("session_id")
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "cookie session_id expired")
			return
		}
		redisSession, err := config.RedisClient.Get(username + "-session_id")
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "cookie session_id illegal")
			return
		}
		if session_id == redisSession {
			c.Next()
			return
		}
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Set cookie failed",
		})
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "username cookie expired")
			return
		}
		userService := service.NewUserService(config.GetMySQLClient())
		user, err := userService.GetUserProfileByName(username)
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "not such user")
			return
		}
		if user.Role != "admin" {
			c.Abort()
			c.String(http.StatusUnauthorized, "not admin")
			return
		}
		c.Next()
	}
}
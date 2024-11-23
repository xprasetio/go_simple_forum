package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/configs"
	"github.com/xprasetio/go_simple_forum.git/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next() // jika kita mempunyai middleware lain, maka middleware ini akan dijalankan sebelum middleware lain
	}
}
func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next() // jika kita mempunyai middleware lain, maka middleware ini akan dijalankan sebelum middleware lain
	}
}

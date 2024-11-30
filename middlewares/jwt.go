package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"data":    nil,
				"message": "Token required, please re login",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"data":    nil,
				"message": "Invalid token, please re login",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, exist := claims["userId"].(float64)
			if !exist {
				c.JSON(http.StatusUnauthorized, gin.H{
					"data":    nil,
					"message": "Invalid payload",
				})
				c.Abort()
				return
			}
			c.Set("userId", uint(userId))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"data":    nil,
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

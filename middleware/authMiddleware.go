package middleware

import (
	"net/http"

	"github.com/HarshitNagpal29/go-jwt-project/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		claims, _, err := helpers.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Next()

	}
}

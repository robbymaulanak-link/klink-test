package middleware

import (
	"net/http"
	"strings"
	pkgjwt "test-k-link-indonesia/packages/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.Set("userLogin", "")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "unauthorized",
			})
			return
		}

		token = strings.Split(token, " ")[1]

		claims, err := pkgjwt.DecodeToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		c.Set("userLogin", claims)
		c.Next()
	}
}

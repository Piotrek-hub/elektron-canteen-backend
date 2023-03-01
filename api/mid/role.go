package mid

import (
	jwtutil "elektron-canteen/foundation/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Role(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Token"] != nil {
			claims, err := jwtutil.DecodeIntoClaims(c.Request.Header["Token"][0])
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			}

			if claims["role"] != requiredRole {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not allowed (role)"})
			}

			c.Next()
		}
	}
}

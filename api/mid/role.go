package mid

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Role(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Authorization"] != nil {
			//t := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
			//claims, err := jwtutil.DecodeIntoClaims(t)
			//if err != nil {
			//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			//}
			if c.Request.Header["role"] != nil && c.Request.Header["role"][0] != requiredRole {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not allowed (role)"})
			}

			c.Next()
		}
	}
}

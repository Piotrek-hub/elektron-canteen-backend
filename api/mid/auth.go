package mid

import (
	"context"
	"elektron-canteen/api/data/user"
	jwtutil "elektron-canteen/foundation/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header["Authorization"] != nil {
			t := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
			token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error while parsing token"})
				}
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			if token.Valid {
				claims, err := jwtutil.DecodeIntoClaims(t)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				}

				objID, err := primitive.ObjectIDFromHex(claims["user"].(string))
				if err != nil {
					panic(err)
				}

				user, err := user.Instance().QueryByID(context.TODO(), objID)
				if err != nil {
					if err == mongo.ErrNoDocuments {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
						return
					}
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					return
				}

				if user.Email != "" {
					c.Request.Header["user_id"] = []string{claims["user"].(string)}
					c.Request.Header["role"] = []string{user.Role}
					c.Next()
				}
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
			return
		}
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func responseWithError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}

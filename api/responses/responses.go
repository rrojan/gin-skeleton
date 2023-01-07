package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func NotFound(c *gin.Context, data string) {
	c.JSON(http.StatusNotFound, gin.H{
		"data": data,
	})
}

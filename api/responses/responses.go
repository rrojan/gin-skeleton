package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func JSON404(c *gin.Context, data interface{}) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": data,
	})
}

func JSONCount(c *gin.Context, data interface{}, count int) {
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"count": count,
	})
}

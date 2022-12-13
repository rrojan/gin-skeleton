package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Writes a StatusOK(200) response to context
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// Writes a NotFound(404) response to context
func JSON404(c *gin.Context, data interface{}) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": data,
	})
}

// Writes a StatusOK(200) response with data and count
func JSONCount(c *gin.Context, data interface{}, count int64) {
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"count": count,
	})
}

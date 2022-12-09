package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rrojan/gin-skeleton/api/responses"
)

type ResourceController struct{}

// GET / -> Lists all resources accessible by user
func (con ResourceController) Index(c *gin.Context) {
	responses.JSON(c, gin.H{
		"data": "This is the resource controller",
	})
}

// GET /:id -> Gets one resource
func (con ResourceController) Get(c *gin.Context) {
	id, _ := c.Params.Get("id")

	responses.JSON(c, gin.H{
		"data": fmt.Sprintf("This is the resource get for %s", id),
	})
}

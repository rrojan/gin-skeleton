package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rrojan/gin-skeleton/api/responses"
)

type ResourceController struct{}

func (con ResourceController) Index(c *gin.Context) {
	responses.JSON(c, gin.H{
		"data": "This is the resource controller",
	})
}

func (con ResourceController) Get(c *gin.Context) {
	id, _ := c.Params.Get("id")

	responses.JSON(c, gin.H{
		"data": fmt.Sprintf("This is the resource get for %s", id),
	})
}

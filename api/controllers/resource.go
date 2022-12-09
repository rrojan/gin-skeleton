package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceController struct{}

func (con ResourceController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "This is the resource controller",
	})
}

func (con ResourceController) Get(c *gin.Context) {
	id, _ := c.Params.Get("id")

	c.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("This is the resource get for %s", id),
	})
}

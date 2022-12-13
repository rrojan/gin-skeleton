package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rrojan/gin-skeleton/api/responses"
	"github.com/rrojan/gin-skeleton/api/services"
)

type ResourceController struct{}

// GET / -> Lists all resources accessible by user
func (con ResourceController) Index(c *gin.Context) {
	data, count := services.ResourceService{}.GetAllResources()

	responses.JSONCount(c, data, count)
}

// GET /:id -> Gets one resource
func (con ResourceController) Get(c *gin.Context) {
	id, _ := c.Params.Get("id")

	responses.JSON(c, fmt.Sprintf("This is the resource get for %s", id))
}

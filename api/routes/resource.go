package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rrojan/gin-skeleton/api/controllers"
)

func ResourceRoute(router *gin.RouterGroup) {
	c := controllers.ResourceController{}
	rg := router.Group("/resources")
	{
		rg.GET("/", c.Index)
		rg.GET("/:id", c.Get)
	}
}

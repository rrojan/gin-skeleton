package routes

import (
	"github.com/gin-gonic/gin"
)

// Router interface
type Router interface{}

// Returns a new router engine instance with routes setup, logging and recovery
func NewRouter() Router {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// V1 Routes
	v1 := router.Group("/v1")
	{
		ResourceRoute(v1)
	}
	return router
}

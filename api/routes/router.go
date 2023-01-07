package routes

import (
	"github.com/gin-gonic/gin"
)

// Router wraps over gin.Engine
type Router struct {
	gin.Engine
	Routes []func(Router)
}

// appRoutes contains all routes to be set up in this app
var appRoutes []func(Router) = []func(Router){
	HealthCheckRoutes,
}

// NewRouter generates a new router instance
func NewRouter() *Router {
	router := gin.New()
	// Use logger for logging and recovery middleware to catch and handle panics
	router.Use(gin.Logger(), gin.Recovery())

	r := &Router{
		*router,
		appRoutes,
	}
	// Bind routes for router
	for _, route := range r.Routes {
		route(*r)
	}

	return r
}

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
	return &Router{
		*gin.Default(),
		appRoutes,
	}
}

// Setup applied routes in appRoutes to application
func (r Router) Setup() {
	for _, route := range r.Routes {
		route(r)
	}
}

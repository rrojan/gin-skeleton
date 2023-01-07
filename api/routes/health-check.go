package routes

import "github.com/rrojan/gin-skeleton/api/controllers"

func HealthCheckRoutes(r Router) {
	c := controllers.NewHealthCheckController()

	r.GET("/health", c.HealthCheck)
}

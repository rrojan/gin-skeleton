package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rrojan/gin-skeleton/api/responses"
)

type HealthCheckController struct{}

// NewHealthCheckController creates a new instance of HealthCheckController
func NewHealthCheckController() HealthCheckController {
	return HealthCheckController{}
}

// HealthCheck returns response of whether Akrit API is live
func (h HealthCheckController) HealthCheck(c *gin.Context) {
	responses.Success(c, "🔥 Gin Skeleton working perfectly! 🔥")
}

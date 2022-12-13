package services

import (
	"github.com/rrojan/gin-skeleton/api/models"
	"github.com/rrojan/gin-skeleton/api/repositories"
)

type ResourceService struct{}

func (s ResourceService) GetAllResources() ([]models.Resource, int64) {
	r := repositories.ResourceRepository{}
	return r.GetAllResources()
}

package repositories

import (
	"github.com/rrojan/gin-skeleton/api/models"
	"github.com/rrojan/gin-skeleton/db"
)

type ResourceRepository struct{}

func (r ResourceRepository) GetAllResources() ([]models.Resource, int64) {
	resources := []models.Resource{}
	var count int64
	db.DB.Find(&resources).Count(&count)
	return resources, count
}

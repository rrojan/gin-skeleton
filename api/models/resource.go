package models

import "gorm.io/gorm"

// TableName: resources
type Resource struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

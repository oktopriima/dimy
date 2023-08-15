package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DeletedAt gorm.DeletedAt
}

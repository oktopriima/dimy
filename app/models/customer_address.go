package models

import (
	"gorm.io/gorm"
)

type CustomerAddress struct {
	ID          int    `json:"id"`
	CustomersID int    `json:"customers_id"`
	Address     string `json:"address"`
	DeletedAt   gorm.DeletedAt
}

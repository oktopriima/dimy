package models

import "gorm.io/gorm"

type PaymentMethod struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	DeletedAt gorm.DeletedAt
}

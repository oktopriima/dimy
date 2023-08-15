package models

import "gorm.io/gorm"

type OrderDetail struct {
	ID         int     `json:"id"`
	OrdersID   int     `json:"orders_id"`
	ProductsID int     `json:"products_id"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
	DeletedAt  gorm.DeletedAt
}

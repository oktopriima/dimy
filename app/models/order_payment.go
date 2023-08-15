package models

import "gorm.io/gorm"

type OrderPayment struct {
	ID               int     `json:"id"`
	OrdersID         int     `json:"orders_id"`
	PaymentMethodsID int     `json:"payment_methods_id"`
	Value            float64 `json:"value"`
	Status           string  `json:"status"`
	DeletedAt        gorm.DeletedAt
}

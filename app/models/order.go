package models

import "gorm.io/gorm"

type Order struct {
	ID                  int     `json:"id"`
	CustomersID         int     `json:"customers_id"`
	CustomerAddressesID int     `json:"customer_addresses_id"`
	TotalPrice          float64 `json:"total_price"`
	Status              string  `json:"status"`
	DeletedAt           gorm.DeletedAt
}

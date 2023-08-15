package seeders

import (
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

type Product struct {
	*Table
}

var products = map[string]float64{
	"coffee":         10000,
	"ice tea":        5000,
	"ice cream cone": 25000,
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{NewTable(db)}
}

func (p *Product) ShouldRun() bool {
	if p.Count("products") > 1 {
		return false
	}
	return true
}

func (p *Product) Run() error {
	for name, price := range products {
		product := new(models.Product)
		product.Name = name
		product.Price = price

		err := p.Create(product)
		if err != nil {
			return err
		}
	}
	return nil
}

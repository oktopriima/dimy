package seeders

import (
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

// CustomersData array name address
var CustomersData = map[string]string{
	"nani": "jl Damai",
	"adi":  "jl Mawar",
	"dodi": "jl halan",
}

type Customers struct {
	*Table
}

func NewCustomers(db *gorm.DB) *Customers {
	return &Customers{NewTable(db)}
}

func (c *Customers) ShouldRun() bool {
	if c.Count("customers") > 1 {
		return false
	}
	return true
}

func (c *Customers) Run() error {
	i := 1
	for name := range CustomersData {
		customer := new(models.Customer)
		customer.ID = i
		customer.Name = name

		err := c.Create(customer)
		if err != nil {
			return err
		}

		i++
	}

	return nil
}

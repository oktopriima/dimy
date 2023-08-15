package seeders

import (
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

type CustomerAddress struct {
	*Table
}

func NewCustomerAddress(db *gorm.DB) *CustomerAddress {
	return &CustomerAddress{NewTable(db)}
}

func (c *CustomerAddress) ShouldRun() bool {
	if c.Count("customer_addresses") > 1 {
		return false
	}
	return true
}
func (c *CustomerAddress) Run() error {
	i := 1
	for _, address := range CustomersData {
		addr := new(models.CustomerAddress)
		addr.CustomersID = i
		addr.Address = address

		err := c.Create(addr)
		if err != nil {
			return err
		}

		i++
	}
	return nil
}

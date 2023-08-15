package seeders

import (
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	*Table
}

func NewPaymentMethod(db *gorm.DB) *PaymentMethod {
	return &PaymentMethod{NewTable(db)}
}

func (p *PaymentMethod) ShouldRun() bool {
	if p.Count("payment_methods") > 1 {
		return false
	}
	return true
}

func (p *PaymentMethod) Run() error {
	payments := []string{
		"BCA", "MANDIRI", "BNI", "OVO", "GOPAY",
	}

	for _, pay := range payments {
		payment := new(models.PaymentMethod)
		payment.Name = pay
		payment.IsActive = true

		err := p.Create(payment)
		if err != nil {
			return err
		}
	}
	return nil
}

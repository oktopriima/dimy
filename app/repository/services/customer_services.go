package services

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"github.com/oktopriima/dimy/app/repository"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Find(ctx context.Context, ID int) (*models.Customer, error) {
	output := new(models.Customer)
	err := c.db.WithContext(ctx).Where("id = ?", ID).Take(&output).Error
	return output, err
}

func (c *customerRepository) Create(ctx context.Context, customer *models.Customer) error {
	err := c.db.WithContext(ctx).Create(customer).Error
	return err
}

func NewCustomerServices(db *gorm.DB) repository.CustomerRepository {
	return &customerRepository{db: db}
}

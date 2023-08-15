package repository

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	Find(ctx context.Context, ID int) (*models.Customer, error)
}

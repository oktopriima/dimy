package repository

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
)

type ProductRepository interface {
	Find(ctx context.Context, ID int) (*models.Product, error)
}

package repository

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(db *gorm.DB, ctx context.Context, order *models.Order) error
	Find(ctx context.Context, ID int) (*models.Order, error)
	Update(db *gorm.DB, ctx context.Context, order *models.Order) error
}

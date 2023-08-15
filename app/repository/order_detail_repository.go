package repository

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Create(db *gorm.DB, ctx context.Context, detail *models.OrderDetail) error
}

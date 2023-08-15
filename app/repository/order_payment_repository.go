package repository

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"gorm.io/gorm"
)

type OrderPaymentRepository interface {
	Create(db *gorm.DB, ctx context.Context, payment *models.OrderPayment) error
	FindRemainingPayment(ctx context.Context, orderID int) ([]*models.OrderPayment, error)
}

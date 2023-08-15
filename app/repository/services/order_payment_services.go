package services

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"github.com/oktopriima/dimy/app/repository"
	"gorm.io/gorm"
)

type orderPaymentRepository struct {
	db *gorm.DB
}

func NewOrderPaymentServices(db *gorm.DB) repository.OrderPaymentRepository {
	return &orderPaymentRepository{db: db}
}

func (o *orderPaymentRepository) Create(db *gorm.DB, ctx context.Context, payment *models.OrderPayment) error {
	return db.WithContext(ctx).Create(payment).Error
}

func (o *orderPaymentRepository) FindRemainingPayment(ctx context.Context, orderID int) ([]*models.OrderPayment, error) {
	var output []*models.OrderPayment
	err := o.db.WithContext(ctx).Where("orders_id = ?", orderID).Find(&output).Error
	return output, err
}

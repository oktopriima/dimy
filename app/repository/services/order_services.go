package services

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"github.com/oktopriima/dimy/app/repository"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderServices(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

func (o *orderRepository) Create(db *gorm.DB, ctx context.Context, order *models.Order) error {
	return db.WithContext(ctx).Create(order).Error
}

func (o *orderRepository) Find(ctx context.Context, ID int) (*models.Order, error) {
	output := new(models.Order)
	err := o.db.WithContext(ctx).Where("id = ?", ID).Take(output).Error
	return output, err
}

func (o *orderRepository) Update(db *gorm.DB, ctx context.Context, order *models.Order) error {
	return db.WithContext(ctx).Save(order).Error
}

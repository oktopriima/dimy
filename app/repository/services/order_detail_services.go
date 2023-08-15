package services

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"github.com/oktopriima/dimy/app/repository"
	"gorm.io/gorm"
)

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailServices(db *gorm.DB) repository.OrderDetailRepository {
	return &orderDetailRepository{db: db}
}

func (o *orderDetailRepository) Create(db *gorm.DB, ctx context.Context, detail *models.OrderDetail) error {
	return db.WithContext(ctx).Create(detail).Error
}

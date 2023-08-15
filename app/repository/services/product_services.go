package services

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"github.com/oktopriima/dimy/app/repository"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Find(ctx context.Context, ID int) (*models.Product, error) {
	output := new(models.Product)
	err := p.db.WithContext(ctx).Where("id = ?", ID).Take(output).Error
	return output, err
}

func NewProductServices(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

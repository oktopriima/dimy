package orders

import (
	"context"
	"github.com/oktopriima/dimy/app/handlers/dto"
	"github.com/oktopriima/dimy/app/repository"
	"gorm.io/gorm"
)

type ordersUseCase struct {
	db               *gorm.DB
	orderRepo        repository.OrderRepository
	orderDetailRepo  repository.OrderDetailRepository
	orderPaymentRepo repository.OrderPaymentRepository
	productRepo      repository.ProductRepository
}

type OrderUseCase interface {
	Create(ctx context.Context, input dto.OrderCreateInput) (*dto.OrderCreateResponse, error)
	Payment(ctx context.Context, input dto.PaymentCreateInput) (*dto.PaymentCreateResponse, error)
}

func NewOrderUseCase(
	db *gorm.DB,
	orderRepo repository.OrderRepository,
	orderDetailRepo repository.OrderDetailRepository,
	orderPaymentRepo repository.OrderPaymentRepository,
	productRepo repository.ProductRepository) OrderUseCase {
	return &ordersUseCase{
		db:               db,
		orderRepo:        orderRepo,
		orderDetailRepo:  orderDetailRepo,
		orderPaymentRepo: orderPaymentRepo,
		productRepo:      productRepo,
	}
}

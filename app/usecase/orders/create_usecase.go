package orders

import (
	"context"
	"github.com/oktopriima/dimy/app/handlers/dto"
	"github.com/oktopriima/dimy/app/helpers/constant"
	"github.com/oktopriima/dimy/app/models"
)

func (o *ordersUseCase) Create(ctx context.Context, input dto.OrderCreateInput) (*dto.OrderCreateResponse, error) {
	var response dto.OrderCreateResponse
	var totalPrice float64
	var orderDetails []*models.OrderDetail
	for _, product := range input.Products {
		existingProduct, err := o.productRepo.Find(ctx, product.ID)
		if err != nil {
			return nil, err
		}

		orderDetails = append(orderDetails, &models.OrderDetail{
			ProductsID: existingProduct.ID,
			Quantity:   product.Quantity,
			Price:      existingProduct.Price,
		})
		totalPrice = totalPrice + (existingProduct.Price * product.Quantity)
	}

	tx := o.db.Begin()
	defer func() {
		tx.Rollback()
	}()

	order := new(models.Order)
	order.CustomersID = input.CustomerID
	order.CustomerAddressesID = input.AddressID
	order.TotalPrice = totalPrice
	order.Status = constant.OrderStatusPending.String()

	if err := o.orderRepo.Create(tx, ctx, order); err != nil {
		return nil, err
	}

	for _, detail := range orderDetails {
		detail.OrdersID = order.ID
		if err := o.orderDetailRepo.Create(tx, ctx, detail); err != nil {
			return nil, err
		}
	}

	tx.Commit()

	response.TotalPrice = totalPrice
	response.OrderCreateInput = input
	response.OrderID = order.ID

	return &response, nil
}

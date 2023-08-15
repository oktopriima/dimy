package orders

import (
	"context"
	"fmt"
	"github.com/oktopriima/dimy/app/handlers/dto"
	"github.com/oktopriima/dimy/app/helpers/constant"
	"github.com/oktopriima/dimy/app/models"
)

func (o *ordersUseCase) Payment(ctx context.Context, input dto.PaymentCreateInput) (*dto.PaymentCreateResponse, error) {
	var output dto.PaymentCreateResponse

	order, err := o.orderRepo.Find(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	if order.Status == constant.OrderStatusSuccess.String() {
		return nil, fmt.Errorf("order already paid")
	}

	currentPayment, err := o.orderPaymentRepo.FindRemainingPayment(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	var totalPaid float64
	for _, pays := range currentPayment {
		totalPaid = totalPaid + pays.Value
	}

	tx := o.db.Begin()
	defer func() {
		tx.Rollback()
	}()

	var remainingPayment float64

	remainingPayment = order.TotalPrice - totalPaid
	for _, payment := range input.Payments {
		orderPayment := new(models.OrderPayment)
		orderPayment.OrdersID = order.ID
		orderPayment.PaymentMethodsID = payment.ID
		orderPayment.Status = constant.PaymentStatusSuccess.String()
		orderPayment.Value = payment.Value

		if err := o.orderPaymentRepo.Create(tx, ctx, orderPayment); err != nil {
			return nil, err
		}

		remainingPayment = remainingPayment - payment.Value

		if remainingPayment < 0 {
			return nil, fmt.Errorf("remaining payment less than 0")
		}
	}

	if remainingPayment == 0 {
		// update order status to success
		order.Status = constant.OrderStatusSuccess.String()
		if err := o.orderRepo.Update(tx, ctx, order); err != nil {
			return nil, err
		}
	}

	tx.Commit()

	output.Order = order
	output.RemainingPayment = remainingPayment
	return &output, nil
}

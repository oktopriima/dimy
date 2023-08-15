package dto

import "github.com/oktopriima/dimy/app/models"

type PaymentCreateInput struct {
	OrderID  int `json:"order_id"`
	Payments []struct {
		ID    int     `json:"id"`
		Value float64 `json:"value"`
	} `json:"payments"`
}

type PaymentCreateResponse struct {
	RemainingPayment float64       `json:"remaining_payment"`
	Order            *models.Order `json:"order"`
}

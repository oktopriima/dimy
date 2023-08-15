package API

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/dimy/app/handlers/dto"
	"github.com/oktopriima/dimy/app/handlers/responses"
	"github.com/oktopriima/dimy/app/usecase/orders"
	"net/http"
)

type OrderHandler struct {
	order orders.OrderUseCase
}

func NewOrderHandler(order orders.OrderUseCase) OrderHandler {
	return OrderHandler{
		order: order,
	}
}

func (h OrderHandler) Create(c echo.Context) error {
	var payload dto.OrderCreateInput
	var response responses.BaseResponse

	ctx := c.Request().Context()

	err := c.Bind(&payload)
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	output, err := h.order.Create(ctx, payload)
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response.Data = output
	response.Success = true
	return c.JSON(http.StatusOK, response)
}

func (h OrderHandler) Payment(c echo.Context) error {
	var payload dto.PaymentCreateInput
	var response responses.BaseResponse

	if err := c.Bind(&payload); err != nil {
		response.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	output, err := h.order.Payment(c.Request().Context(), payload)
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response.Success = true
	response.Data = output
	return c.JSON(http.StatusOK, response)
}

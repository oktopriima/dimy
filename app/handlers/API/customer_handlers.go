package API

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/dimy/app/handlers/responses"
	"github.com/oktopriima/dimy/app/usecase/customers"
	"net/http"
	"strconv"
)

type CustomerHandlers struct {
	customer customers.CustomerUseCase
}

func NewCustomerHandlers(customer customers.CustomerUseCase) CustomerHandlers {
	return CustomerHandlers{
		customer: customer,
	}
}

func (h CustomerHandlers) Find(c echo.Context) error {
	var response responses.BaseResponse
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	output, err := h.customer.FindUseCase(ctx, id)
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response.Data = output
	response.Success = true
	return c.JSON(http.StatusOK, response)
}

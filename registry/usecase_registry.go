package registry

import (
	"github.com/oktopriima/dimy/app/usecase/customers"
	"github.com/oktopriima/dimy/app/usecase/orders"
	"go.uber.org/dig"
)

func NewUseCaseRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(customers.NewCustomerUseCase); err != nil {
		panic(err)
	}

	if err = container.Provide(orders.NewOrderUseCase); err != nil {
		panic(err)
	}

	return container
}

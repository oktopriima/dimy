package registry

import (
	"github.com/oktopriima/dimy/app/repository/services"
	"go.uber.org/dig"
)

func NewRepositoryRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(services.NewCustomerServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewOrderServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewOrderDetailServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewOrderPaymentServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewProductServices); err != nil {
		panic(err)
	}

	return container
}

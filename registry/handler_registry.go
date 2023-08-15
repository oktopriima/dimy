package registry

import (
	"github.com/oktopriima/dimy/app/handlers/API"
	"go.uber.org/dig"
)

func NewHandlerRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(API.NewCustomerHandlers); err != nil {
		panic(err)
	}

	if err = container.Provide(API.NewOrderHandler); err != nil {
		panic(err)
	}

	return container
}

package registry

import (
	"github.com/oktopriima/dimy/app/router"
	"go.uber.org/dig"
)

func NewRouterRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Invoke(router.NewAPIRouter); err != nil {
		panic(err)
	}

	return container
}

package main

import (
	"github.com/oktopriima/dimy/registry"
	"go.uber.org/dig"
)

func NewInject() *dig.Container {
	c := dig.New()

	c = registry.NewConfigRegistry(c)
	c = registry.NewRepositoryRegistry(c)
	c = registry.NewUseCaseRegistry(c)
	c = registry.NewHandlerRegistry(c)

	c = registry.NewRouterRegistry(c)

	return c
}

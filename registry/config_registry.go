package registry

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/dimy/config"
	"go.uber.org/dig"
)

func NewConfigRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(config.NewEnvironment); err != nil {
		panic(err)
	}

	if err = container.Provide(func() *echo.Echo {
		return echo.New()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(config.NewMysqlConnector); err != nil {
		panic(err)
	}

	if err = container.Provide(config.NewServerInstance); err != nil {
		panic(err)
	}

	if err = container.Provide(config.NewSeedersInstance); err != nil {
		panic(err)
	}

	return container
}

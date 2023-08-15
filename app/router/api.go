package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/oktopriima/dimy/app/handlers/API"
	"net/http"
)

func NewAPIRouter(
	e *echo.Echo,
	customer API.CustomerHandlers,
	order API.OrderHandler,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:       middleware.DefaultSkipper,
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Authorization"},
	}))

	route := e.Group("api")

	route.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, struct {
			Message string
		}{Message: "all is good"})
	})

	{
		customerRoute := route.Group("/customers")
		customerRoute.GET("/:id", customer.Find)
	}

	{
		orderRoute := route.Group("/orders")
		orderRoute.POST("", order.Create)
		orderRoute.POST("/payment", order.Payment)
	}

}

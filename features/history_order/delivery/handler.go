package delivery

import (
	"project/kutsuya/features/history_order"

	"github.com/labstack/echo/v4"
)

type HistoryDelivery struct {
	// historyUsecase history_order.UsecaseInterface
}

func New(e *echo.Echo, usecase history_order.UsecaseInterface) {
	// handler := &HistoryDelivery{
	// 	historyUsecase: usecase,
	// }

	// e.POST("/carts", handler.PostNewCart, middlewares.JWTMiddleware())
	// e.GET("/carts", handler.GetAllCarts, middlewares.JWTMiddleware())
}

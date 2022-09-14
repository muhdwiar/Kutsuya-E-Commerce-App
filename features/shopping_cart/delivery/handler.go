package delivery

import (
	"project/kutsuya/features/shopping_cart"

	"github.com/labstack/echo/v4"
)

type CartDelivery struct {
	cartUsecase shopping_cart.UsecaseInterface
}

func New(e *echo.Echo, usecase shopping_cart.UsecaseInterface) {
	// handler := &UserDelivery{
	// 	userUsecase: usecase,
	// }
}

package delivery

import (
	"project/kutsuya/features/user"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	// handler := &UserDelivery{
	// 	userUsecase: usecase,
	// }
}

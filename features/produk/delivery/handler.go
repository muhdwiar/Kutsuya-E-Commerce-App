package delivery

import (
	"project/kutsuya/features/produk"

	"github.com/labstack/echo/v4"
)

type ProdukDelivery struct {
	produkUsecase produk.UsecaseInterface
}

func New(e *echo.Echo, usecase produk.UsecaseInterface) {
	// handler := &ProdukDelivery{
	// 	produkUsecase: usecase,
	// }
}

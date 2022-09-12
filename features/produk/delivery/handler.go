package delivery

import (
	"net/http"
	"project/kutsuya/features/produk"
	"project/kutsuya/utils/helper"

	"github.com/labstack/echo/v4"
)

type ProdukDelivery struct {
	produkUsecase produk.UsecaseInterface
}

func New(e *echo.Echo, usecase produk.UsecaseInterface) {
	handler := &ProdukDelivery{
		produkUsecase: usecase,
	}

	e.GET("/products", handler.Get_AllProduk)
}

func (delivery *ProdukDelivery) Get_AllProduk(c echo.Context) error {
	result, err := delivery.produkUsecase.Get_AllProduk()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get all product data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all products data", FromCoreList(result)))

}

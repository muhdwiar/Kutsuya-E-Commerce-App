package delivery

import (
	"net/http"
	"project/kutsuya/features/produk"
	"project/kutsuya/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProdukDelivery struct {
	produkUsecase produk.UsecaseInterface
}

func New(e *echo.Echo, usecase produk.UsecaseInterface) {
	handler := &ProdukDelivery{
		produkUsecase: usecase,
	}

	e.PUT("/products/:id", handler.Put_ProdukData)
}

func (delivery *ProdukDelivery) Put_ProdukData(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	var dataRequest ProdukRequest
	errBind := c.Bind(&dataRequest)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind produk data"))
	}

	row, err := delivery.produkUsecase.PutProduk(ToCore(dataRequest, id_conv))

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update produk data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("succes update produk data"))

}

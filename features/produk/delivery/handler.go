package delivery

import (
	"net/http"
	"project/kutsuya/features/produk"
	"project/kutsuya/middlewares"
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
	e.POST("/products", handler.PostProduk, middlewares.JWTMiddleware())
}

func (delivery *ProdukDelivery) Get_AllProduk(c echo.Context) error {
	result, err := delivery.produkUsecase.Get_AllProduk()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get all product data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all products data", FromCoreList(result)))

}

func (delivery *ProdukDelivery) PostProduk(c echo.Context) error {
	var produk_RequestData ProdukRequest
	errBind := c.Bind(&produk_RequestData)

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}

	produk_RequestData.User_Id = uint(userId)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind produk data"))
	}

	row, err := delivery.produkUsecase.PostProduk(ToCore(produk_RequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input produk data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input produk data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input new produk"))

}

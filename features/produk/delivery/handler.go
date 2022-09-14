package delivery

import (
	"net/http"
	"project/kutsuya/features/produk"
	"project/kutsuya/middlewares"
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

	e.GET("/products", handler.Get_AllProduk)
	e.POST("/products", handler.PostProduk, middlewares.JWTMiddleware())
	e.PUT("/products/:id", handler.Put_ProdukData, middlewares.JWTMiddleware())
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

	row, err := delivery.produkUsecase.PutProduk(ToCore(dataRequest), id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update produk data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("succes update produk data"))

}

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
	e.GET("/products/:id", handler.Get_ProdukById)
	e.DELETE("/products/:id", handler.Delete_ProdukData, middlewares.JWTMiddleware())
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
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind produk data"))
	}

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}
	produk_RequestData.User_Id = uint(userId)

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

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}

	var produk_RequestData ProdukRequest
	errBind := c.Bind(&produk_RequestData)
	produk_RequestData.User_Id = uint(userId)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind produk data"))
	}

	row, err := delivery.produkUsecase.PutProduk(ToCore(produk_RequestData), id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update produk data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("succes update produk data"))

}

func (delivery *ProdukDelivery) Get_ProdukById(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.produkUsecase.GetProdukById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get product data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get product data", FromCore(result)))

}

func (delivery *ProdukDelivery) Delete_ProdukData(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Decrypt Jwt Token"))
	}

	row, err := delivery.produkUsecase.DeleteProdukUser(userId, id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Delete Produk Data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Delete Produk Data"))

}

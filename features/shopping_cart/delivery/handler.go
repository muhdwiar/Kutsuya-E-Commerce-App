package delivery

import (
	"net/http"
	"project/kutsuya/features/shopping_cart"
	"project/kutsuya/middlewares"
	"project/kutsuya/utils/helper"

	"github.com/labstack/echo/v4"
)

type CartDelivery struct {
	cartUsecase shopping_cart.UsecaseInterface
}

func New(e *echo.Echo, usecase shopping_cart.UsecaseInterface) {
	handler := &CartDelivery{
		cartUsecase: usecase,
	}

	e.POST("/carts", handler.PostNewCart, middlewares.JWTMiddleware())
	e.GET("/carts", handler.GetAllCarts, middlewares.JWTMiddleware())
}

func (delivery *CartDelivery) PostNewCart(c echo.Context) error {
	var cart_RequestData CartRequest
	errBind := c.Bind(&cart_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind Data"))
	}

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Decrypt Jwt Token"))
	}
	cart_RequestData.User_Id = uint(userId)

	row, err := delivery.cartUsecase.InsertCart(ToCore(cart_RequestData))

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Make New Cart"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Make New Cart"))
}

func (delivery *CartDelivery) GetAllCarts(c echo.Context) error {

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Decrypt Jwt Token"))
	}

	result, err := delivery.cartUsecase.SelectCarts(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Get All Cart Data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("Success Get All Cart Data", FromCoreList(result)))

}

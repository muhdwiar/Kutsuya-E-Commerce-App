package delivery

import (
	"net/http"
	"project/kutsuya/features/history_order"
	"project/kutsuya/middlewares"
	"project/kutsuya/utils/helper"

	"github.com/labstack/echo/v4"
)

type HistoryDelivery struct {
	historyUsecase history_order.UsecaseInterface
}

func New(e *echo.Echo, usecase history_order.UsecaseInterface) {
	handler := &HistoryDelivery{
		historyUsecase: usecase,
	}

	e.POST("/orders", handler.PostHistoryOrder, middlewares.JWTMiddleware())
	e.POST("/cancel", handler.PostHistoryCancel, middlewares.JWTMiddleware())
	e.GET("/orders", handler.GetAllOrders, middlewares.JWTMiddleware())

}

func (delivery *HistoryDelivery) PostHistoryOrder(c echo.Context) error {
	var historyList_RequestData []HistoryRequest
	errBind := c.Bind(&historyList_RequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind Data"))
	}

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Decrypt Jwt Token"))
	}

	row, err := delivery.historyUsecase.InsertHistoryOrder(ToCoreRequestList(historyList_RequestData), userId)

	if err != nil || row != len(historyList_RequestData) {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Input History Order"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Input All History Order"))

}

func (delivery *HistoryDelivery) PostHistoryCancel(c echo.Context) error {
	var cancel_RequestData HistoryRequest
	errBind := c.Bind(&cancel_RequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind Data"))
	}

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Decrypt Jwt Token"))
	}

	row, err := delivery.historyUsecase.InsertHistoryCancel(ToCoreRequest(cancel_RequestData), userId)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Cancel Order"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Cancel Order"))

}

func (delivery *HistoryDelivery) GetAllOrders(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Decrypt Jwt Token"))
	}

	result, err := delivery.historyUsecase.SelectOrders(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Get All Orders"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("Success Get All Orders", FromCoreList(result)))

}

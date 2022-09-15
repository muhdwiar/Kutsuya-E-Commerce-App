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
	// e.GET("/carts", handler.GetAllCarts, middlewares.JWTMiddleware())
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
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Input All History Order"))

}

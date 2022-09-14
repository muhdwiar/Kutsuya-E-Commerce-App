package delivery

import (
	"net/http"
	"project/kutsuya/features/user"
	"project/kutsuya/middlewares"
	"project/kutsuya/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	// e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/login", handler.LoginUser)
	e.POST("/users", handler.PostData)
	e.GET("/users/:id", handler.GetUserById, middlewares.JWTMiddleware())

}

func (delivery *UserDelivery) GetUserById(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Conv Id"))
	}
	result, err := delivery.userUsecase.GetById(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Get User Data By Id"))
	}
	return c.JSON(http.StatusOK, helper.Success_DataResp("Success Get Data By Id", result))

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var userRequestData UserRequest
	errBind := c.Bind(&userRequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind User Data"))
	}

	token, row, err := delivery.userUsecase.PostData(ToCore(userRequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Input User Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Insert Row Affected Is Not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("Success Insert", token))

}

func (delivery *UserDelivery) LoginUser(c echo.Context) error {
	var userRequest_Login UserRequest
	errBind := c.Bind(&userRequest_Login)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Data Doesn't Exist"))
	}

	Token_JWT, err := delivery.userUsecase.PostLogin(ToCore(userRequest_Login))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Data Doesn't Exist"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("Success Login", Token_JWT))

}

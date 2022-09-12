package factory

import (
	userData "project/kutsuya/features/user/data"
	userDelivery "project/kutsuya/features/user/delivery"
	userUsecase "project/kutsuya/features/user/usecase"

	produkData "project/kutsuya/features/produk/data"
	produkDelivery "project/kutsuya/features/produk/delivery"
	produkUsecase "project/kutsuya/features/produk/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	produkDataFactory := produkData.New(db)
	produkUsecaseFactory := produkUsecase.New(produkDataFactory)
	produkDelivery.New(e, produkUsecaseFactory)

}

package migration

import (
	historyModel "project/kutsuya/features/history_order/data"
	produkModel "project/kutsuya/features/produk/data"
	cartModel "project/kutsuya/features/shopping_cart/data"
	userModel "project/kutsuya/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&produkModel.Produk{})
	db.AutoMigrate(&cartModel.Shopping_Cart{})
	db.AutoMigrate(&historyModel.History_Order{})
}

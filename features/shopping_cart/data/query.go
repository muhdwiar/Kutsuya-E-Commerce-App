package data

import (
	"project/kutsuya/features/shopping_cart"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) shopping_cart.DataInterface {
	return &cartData{
		db: db,
	}

}

package data

import (
	"project/kutsuya/features/produk"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) produk.DataInterface {
	return &productData{
		db: db,
	}

}

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

func (repo *productData) Select_AllProduk() ([]produk.Core, error) {
	var all_ProdData []Produk
	tx := repo.db.Find(&all_ProdData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	produk_List := toCoreList(all_ProdData)
	return produk_List, nil
}

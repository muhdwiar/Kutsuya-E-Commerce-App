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

func (repo *productData) InsertProduk(newProduk produk.Core) (int, error) {
	newUser := fromCore(newProduk)

	tx := repo.db.Create(&newUser)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

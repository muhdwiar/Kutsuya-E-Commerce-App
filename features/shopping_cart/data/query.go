package data

import (
	cart "project/kutsuya/features/shopping_cart"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.DataInterface {
	return &cartData{
		db: db,
	}

}

func (repo *cartData) CreateCart(dataCart cart.Core) (int, error) {
	newCart := fromCore(dataCart)

	tx := repo.db.Create(&newCart)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *cartData) FindCarts(user_id int) ([]cart.Core, error) {
	var all_CartData []Shopping_Cart
	tx := repo.db.Where("user_id = ?", user_id).Find(&all_CartData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	produk_List := toCoreList(all_CartData)
	return produk_List, nil

}

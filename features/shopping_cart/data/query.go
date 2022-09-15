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

	cart_List := toCoreList(all_CartData)
	return cart_List, nil

}

func (repo *cartData) GetCartByID(cart_id int) (cart.Core, error) {
	var CartData Shopping_Cart
	tx := repo.db.Find(&CartData, cart_id)

	if tx.Error != nil {
		return cart.Core{}, tx.Error
	}

	return CartData.toCore(), nil

}

func (repo *cartData) DelCartByID(cart_id, user_id int) (int, error) {
	var modelCart Shopping_Cart
	tx := repo.db.Where("user_id = ?", user_id).Delete(&modelCart, cart_id)

	if tx.Error != nil {
		return -1, tx.Error
	}

	return int(tx.RowsAffected), nil

}

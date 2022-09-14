package data

import (
	produkModel "project/kutsuya/features/produk/data"
	"project/kutsuya/features/shopping_cart"
	userModel "project/kutsuya/features/user/data"

	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	Jumlah      int
	Total_Biaya int
	User_Id     uint
	User        userModel.User
	Product_Id  uint
	Produk      produkModel.Produk
}

func fromCore(dataCore shopping_cart.Core) Shopping_Cart {
	return Shopping_Cart{
		Jumlah:      dataCore.Jumlah,
		Total_Biaya: dataCore.Total_Biaya,
		User_Id:     dataCore.User_Id,
		Product_Id:  dataCore.Product_Id,
	}

}

func (dataCart *Shopping_Cart) toCore() shopping_cart.Core {
	return shopping_cart.Core{
		ID:          dataCart.ID,
		Jumlah:      dataCart.Jumlah,
		Total_Biaya: dataCart.Total_Biaya,
		User_Id:     dataCart.User_Id,
		Product_Id:  dataCart.Product_Id,
	}
}

func toCoreList(dataCart []Shopping_Cart) []shopping_cart.Core {
	var dataCore []shopping_cart.Core

	for key := range dataCart {
		dataCore = append(dataCore, dataCart[key].toCore())

	}

	return dataCore

}

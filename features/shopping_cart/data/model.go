package data

import (
	produkModel "project/kutsuya/features/produk/data"
	"project/kutsuya/features/shopping_cart"
	userModel "project/kutsuya/features/user/data"

	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	User_Id     uint
	Product_Id  uint
	Nama_Produk string
	Ukuran      int
	Merk        string
	Biaya       int
	File_Image  string
	User        userModel.User     `gorm:"foreignKey:User_Id"`
	Produk      produkModel.Produk `gorm:"foreignKey:Product_Id"`
}

func fromCore(dataCore shopping_cart.Core) Shopping_Cart {
	return Shopping_Cart{
		User_Id:     dataCore.User_Id,
		Product_Id:  dataCore.Product_Id,
		Nama_Produk: dataCore.Nama_Produk,
		Ukuran:      dataCore.Ukuran,
		Merk:        dataCore.Merk,
		Biaya:       dataCore.Biaya,
		File_Image:  dataCore.File_Image,
	}

}

func (dataCart *Shopping_Cart) toCore() shopping_cart.Core {
	return shopping_cart.Core{
		ID:          dataCart.ID,
		User_Id:     dataCart.User_Id,
		Product_Id:  dataCart.Product_Id,
		Nama_Produk: dataCart.Nama_Produk,
		Ukuran:      dataCart.Ukuran,
		Merk:        dataCart.Merk,
		Biaya:       dataCart.Biaya,
		File_Image:  dataCart.File_Image,
	}
}

func toCoreList(dataCart []Shopping_Cart) []shopping_cart.Core {
	var dataCore []shopping_cart.Core

	for key := range dataCart {
		dataCore = append(dataCore, dataCart[key].toCore())

	}

	return dataCore

}

package data

import (
	// produkModel "project/kutsuya/features/produk/data"
	"project/kutsuya/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama_User     string
	Email         string
	Password      string
	Alamat        string
	Produk        []Produk
	Shopping_Cart []Shopping_Cart
}

type Produk struct {
	gorm.Model
	Nama_Produk     string
	Ukuran          int
	Merk            string
	Warna           string
	Gender_Pengguna string
	Harga           int
	Desksripsi      string
	File_Image      string
	User_Id         uint
	User            User
}

type Shopping_Cart struct {
	gorm.Model
	Jumlah      int
	Total_Biaya int
	User_Id     uint
	User        User
	Product_Id  uint
	Produk      Produk
}

func fromCore(dataCore user.Core) User {
	return User{
		Nama_User: dataCore.Nama_User,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Alamat:    dataCore.Alamat,
	}

}

func (dataUser *User) toCore() user.Core {
	return user.Core{
		ID:        dataUser.ID,
		Nama_User: dataUser.Nama_User,
		Email:     dataUser.Email,
		Password:  dataUser.Password,
		Alamat:    dataUser.Alamat,
	}
}

func toCoreList(dataUser []User) []user.Core {
	var dataCore []user.Core

	for key := range dataUser {
		dataCore = append(dataCore, dataUser[key].toCore())

	}

	return dataCore

}

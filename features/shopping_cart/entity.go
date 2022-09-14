package shopping_cart

import (
	"project/kutsuya/features/produk"
	"project/kutsuya/features/user"
	"time"
)

type Core struct {
	ID          uint
	User_Id     uint
	Product_Id  uint
	Nama_Produk string
	Ukuran      int
	Merk        string
	Biaya       int
	File_Image  string
	User        user.Core
	Produk      produk.Core
	Created_At  time.Time
}

type UsecaseInterface interface {
}

type DataInterface interface {
}

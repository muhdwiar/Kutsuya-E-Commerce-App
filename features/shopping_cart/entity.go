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
	Harga       int
	File_Image  string
	User        user.Core
	Produk      produk.Core
	Created_At  time.Time
}

type UsecaseInterface interface {
	InsertCart(data Core) (row int, err error)
	SelectCarts(user_id int) (data []Core, err error)
}

type DataInterface interface {
	CreateCart(data Core) (row int, err error)
	FindCarts(user_id int) (data []Core, err error)
	GetCartByID(cart_id int) (data Core, err error)
	DelCartByID(cart_id, user_id int) (row int, err error)
}

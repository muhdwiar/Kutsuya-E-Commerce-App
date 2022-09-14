package shopping_cart

import (
	"project/kutsuya/features/produk"
	"project/kutsuya/features/user"
	"time"
)

type Core struct {
	ID          uint
	Jumlah      int
	Total_Biaya int
	User_Id     uint
	User        user.Core
	Product_Id  uint
	Produk      produk.Core
	Created_At  time.Time
	Updated_At  time.Time
}

type UsecaseInterface interface {
}

type DataInterface interface {
}

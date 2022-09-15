package history_order

import (
	user "project/kutsuya/features/user"
	"time"
)

type Core struct {
	ID           uint
	User_Id      uint
	Nama_Produk  string
	Jumlah       int
	Total_Biaya  int
	Status_Order string
	User         user.Core
	Created_At   time.Time
}

type CoreRequest struct {
	ID         uint
	User_Id    uint
	Cart_Id    uint
	Jumlah     int
	User       user.Core
	Created_At time.Time
}

type UsecaseInterface interface {
	InsertHistoryOrder(data []CoreRequest, user_id int) (row int, err error)
}

type DataInterface interface {
	CreateHistoryOrder(data Core) (row int, err error)
}

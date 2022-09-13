package produk

import (
	"project/kutsuya/features/user"
	"time"
)

type Core struct {
	ID              uint
	Nama_Produk     string
	Ukuran          int
	Merk            string
	Warna           string
	Gender_Pengguna string
	Harga           int
	Desksripsi      string
	File_Image      string
	Created_At      time.Time
	Updated_At      time.Time
	User_Id         uint
	User            user.Core
}

type UsecaseInterface interface {
	Get_AllProduk() (data []Core, err error)
	PostProduk(data Core) (row int, err error)
	PutProduk(data Core, id int) (row int, err error)
}

type DataInterface interface {
	Select_AllProduk() (data []Core, err error)
	InsertProduk(data Core) (row int, err error)
	UpdateDataProduk(data Core, id int) (row int, err error)
}

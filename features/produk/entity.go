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
	User            user.Core
}

type UsecaseInterface interface {
	PutProduk(data Core) (row int, err error)
}

type DataInterface interface {
	UpdateDataProduk(data Core) (row int, err error)
}

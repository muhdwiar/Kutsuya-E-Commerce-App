package delivery

import "project/kutsuya/features/produk"

type ProdukRequest struct {
	Nama_Produk     string `json:"nama_produk" form:"nama_produk"`
	Ukuran          int    `json:"ukuran" form:"ukuran"`
	Merk            string `json:"merk" form:"merk"`
	Warna           string `json:"warna" form:"warna"`
	Gender_Pengguna string `json:"gender_pengguna" form:"gender_pengguna"`
	Harga           int    `json:"harga" form:"harga"`
	Desksripsi      string `json:"deskripsi" form:"deskripsi"`
	File_Image      string `json:"file_image" form:"file_image"`
	User_Id         uint   `json:"user_id" form:"user_id"`
}

func ToCore(dataRequest ProdukRequest) produk.Core {
	return produk.Core{
		Nama_Produk:     dataRequest.Nama_Produk,
		Ukuran:          dataRequest.Ukuran,
		Merk:            dataRequest.Merk,
		Warna:           dataRequest.Warna,
		Gender_Pengguna: dataRequest.Gender_Pengguna,
		Harga:           dataRequest.Harga,
		Desksripsi:      dataRequest.Desksripsi,
		File_Image:      dataRequest.File_Image,
		User_Id:         dataRequest.User_Id,
	}

}

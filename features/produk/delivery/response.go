package delivery

import "project/kutsuya/features/produk"

type ProdukResponse struct {
	ID              uint   `json:"id" form:"id"`
	Nama_Produk     string `json:"nama_produk" form:"nama_produk"`
	Ukuran          int    `json:"ukuran" form:"ukuran"`
	Merk            string `json:"merk" form:"merk"`
	Warna           string `json:"warna" form:"warna"`
	Gender_Pengguna string `json:"gender_pengguna" form:"gender_pengguna"`
	Harga           int    `json:"harga" form:"harga"`
	Desksripsi      string `json:"deskripsi" form:"deskripsi"`
	File_Image      string `json:"file_image" form:"file_image"`
}

func FromCore(dataCore produk.Core) ProdukResponse {
	return ProdukResponse{
		ID:              dataCore.ID,
		Nama_Produk:     dataCore.Nama_Produk,
		Ukuran:          dataCore.Ukuran,
		Merk:            dataCore.Merk,
		Warna:           dataCore.Warna,
		Gender_Pengguna: dataCore.Gender_Pengguna,
		Harga:           dataCore.Harga,
		Desksripsi:      dataCore.Desksripsi,
		File_Image:      dataCore.File_Image,
	}

}

func FromCoreList(dataCore []produk.Core) []ProdukResponse {
	var dataResponse []ProdukResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}

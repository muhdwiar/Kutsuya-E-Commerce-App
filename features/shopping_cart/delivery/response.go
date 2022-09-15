package delivery

import (
	"project/kutsuya/features/shopping_cart"
)

type CartResponse struct {
	ID          uint   `json:"id" form:"id"`
	User_Id     uint   `json:"user_id" form:"user_id"`
	Product_Id  uint   `json:"product_id" form:"product_id"`
	Nama_Produk string `json:"nama_produk" form:"nama_produk"`
	Ukuran      int    `json:"ukuran" form:"ukuran"`
	Merk        string `json:"merk" form:"merk"`
	Harga       int    `json:"harga" form:"harga"`
	File_Image  string `json:"file_image" form:"file_image"`
}

func FromCore(dataCore shopping_cart.Core) CartResponse {
	return CartResponse{
		ID:          dataCore.ID,
		User_Id:     dataCore.User_Id,
		Product_Id:  dataCore.Product_Id,
		Nama_Produk: dataCore.Nama_Produk,
		Ukuran:      dataCore.Ukuran,
		Merk:        dataCore.Merk,
		Harga:       dataCore.Harga,
		File_Image:  dataCore.File_Image,
	}

}

func FromCoreList(dataCore []shopping_cart.Core) []CartResponse {
	var dataResponse []CartResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}

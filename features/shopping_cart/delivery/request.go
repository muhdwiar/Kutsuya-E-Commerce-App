package delivery

import (
	"project/kutsuya/features/shopping_cart"
)

type CartRequest struct {
	User_Id     uint   `json:"user_id" form:"user_id"`
	Product_Id  uint   `json:"product_id" form:"product_id"`
	Nama_Produk string `json:"nama_produk" form:"nama_produk"`
	Ukuran      int    `json:"ukuran" form:"ukuran"`
	Merk        string `json:"merk" form:"merk"`
	Biaya       int    `json:"biaya" form:"biaya"`
	File_Image  string `json:"file_image" form:"file_image"`
}

func ToCore(dataRequest CartRequest) shopping_cart.Core {
	return shopping_cart.Core{
		User_Id:     dataRequest.User_Id,
		Product_Id:  dataRequest.Product_Id,
		Nama_Produk: dataRequest.Nama_Produk,
		Ukuran:      dataRequest.Ukuran,
		Merk:        dataRequest.Merk,
		Biaya:       dataRequest.Biaya,
		File_Image:  dataRequest.File_Image,
	}

}

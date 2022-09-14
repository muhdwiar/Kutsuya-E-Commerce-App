package delivery

import (
	"project/kutsuya/features/shopping_cart"
)

type CartRequest struct {
	Jumlah      int  `json:"jumlah" form:"jumlah"`
	Total_Biaya int  `json:"total_biaya" form:"total_biaya"`
	User_Id     uint `json:"user_id" form:"user_id"`
	Product_Id  uint `json:"product_id" form:"product_id"`
}

func ToCore(dataRequest CartRequest) shopping_cart.Core {
	return shopping_cart.Core{
		Jumlah:      dataRequest.Jumlah,
		Total_Biaya: dataRequest.Total_Biaya,
		User_Id:     dataRequest.User_Id,
		Product_Id:  dataRequest.Product_Id,
	}

}

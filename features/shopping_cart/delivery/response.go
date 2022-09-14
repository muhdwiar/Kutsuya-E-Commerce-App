package delivery

import (
	"project/kutsuya/features/shopping_cart"
)

type CartResponse struct {
	ID          uint `json:"id" form:"id"`
	Jumlah      int  `json:"jumlah" form:"jumlah"`
	Total_Biaya int  `json:"total_biaya" form:"total_biaya"`
	User_Id     uint `json:"user_id" form:"user_id"`
	Product_Id  uint `json:"product_id" form:"product_id"`
}

func FromCore(dataCore shopping_cart.Core) CartResponse {
	return CartResponse{
		ID:          dataCore.ID,
		Jumlah:      dataCore.Jumlah,
		Total_Biaya: dataCore.Total_Biaya,
		User_Id:     dataCore.User_Id,
		Product_Id:  dataCore.Product_Id,
	}

}

func FromCoreList(dataCore []shopping_cart.Core) []CartResponse {
	var dataResponse []CartResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}

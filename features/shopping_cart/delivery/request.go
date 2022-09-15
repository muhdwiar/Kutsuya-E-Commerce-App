package delivery

import (
	"project/kutsuya/features/shopping_cart"
)

type CartRequest struct {
	User_Id    uint `json:"user_id" form:"user_id"`
	Product_Id uint `json:"product_id" form:"product_id"`
}

func ToCore(dataRequest CartRequest) shopping_cart.Core {
	return shopping_cart.Core{
		User_Id:    dataRequest.User_Id,
		Product_Id: dataRequest.Product_Id,
	}

}

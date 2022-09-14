package usecase

import (
	"project/kutsuya/features/shopping_cart"
)

type cartUsecase struct {
	cartData shopping_cart.DataInterface
}

func New(data shopping_cart.DataInterface) shopping_cart.UsecaseInterface {
	return &cartUsecase{
		cartData: data,
	}
}

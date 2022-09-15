package usecase

import (
	"errors"
	cart "project/kutsuya/features/shopping_cart"
)

type cartUsecase struct {
	cartData cart.DataInterface
}

func New(data cart.DataInterface) cart.UsecaseInterface {
	return &cartUsecase{
		cartData: data,
	}
}

func (usecase *cartUsecase) InsertCart(dataCart cart.Core) (int, error) {
	if dataCart.Product_Id == 0 || dataCart.User_Id == 0 {
		return -1, errors.New("data Product_Id and User_Id is null")

	}

	row, err := usecase.cartData.CreateCart(dataCart)
	return row, err

}

func (usecase *cartUsecase) SelectCarts(user_id int) ([]cart.Core, error) {
	results, err := usecase.cartData.FindCarts(user_id)
	return results, err

}

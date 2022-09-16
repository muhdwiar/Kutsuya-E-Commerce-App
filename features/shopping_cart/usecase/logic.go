package usecase

import (
	"errors"
	"project/kutsuya/features/produk"
	cart "project/kutsuya/features/shopping_cart"
)

type cartUsecase struct {
	cartData   cart.DataInterface
	produkData produk.DataInterface
}

func New(dataCart cart.DataInterface, dataProduk produk.DataInterface) cart.UsecaseInterface {
	return &cartUsecase{
		cartData:   dataCart,
		produkData: dataProduk,
	}
}

func (usecase *cartUsecase) InsertCart(dataCart cart.Core) (int, error) {

	var newCart cart.Core

	if dataCart.Product_Id == 0 || dataCart.User_Id == 0 {
		return -1, errors.New("data Product_Id dan User_Id kosong")
	}

	dataProduk, errProduk := usecase.produkData.SelectProdukById(int(dataCart.Product_Id))
	if errProduk != nil {
		return -1, errProduk
	}

	newCart.User_Id = dataCart.User_Id
	newCart.Product_Id = dataCart.Product_Id
	newCart.Nama_Produk = dataProduk.Nama_Produk
	newCart.Ukuran = dataProduk.Ukuran
	newCart.Merk = dataProduk.Merk
	newCart.Harga = dataProduk.Harga
	newCart.File_Image = dataProduk.File_Image

	row, err := usecase.cartData.CreateCart(newCart)
	return row, err

}

func (usecase *cartUsecase) SelectCarts(user_id int) ([]cart.Core, error) {
	results, err := usecase.cartData.FindCarts(user_id)
	return results, err

}

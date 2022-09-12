package usecase

import (
	"project/kutsuya/features/produk"
)

type produkUsecase struct {
	produkData produk.DataInterface
}

func New(data produk.DataInterface) produk.UsecaseInterface {
	return &produkUsecase{
		produkData: data,
	}

}

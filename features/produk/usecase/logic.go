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

func (usecase *produkUsecase) Get_AllProduk() ([]produk.Core, error) {
	results, err := usecase.produkData.Select_AllProduk()
	return results, err

}

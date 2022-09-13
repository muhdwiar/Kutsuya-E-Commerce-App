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

func (usecase *produkUsecase) PutProduk(data produk.Core) (int, error) {
	row, err := usecase.produkData.UpdateDataProduk(data)
	return row, err

}

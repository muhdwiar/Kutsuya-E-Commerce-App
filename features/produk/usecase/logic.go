package usecase

import (
	"errors"
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

func (usecase *produkUsecase) PostProduk(newProduk produk.Core) (int, error) {
	if newProduk.Nama_Produk == "" || newProduk.Harga == 0 {
		return -1, errors.New("nama produk dan harga tidak boleh dikosongkan")

	}

	row, err := usecase.produkData.InsertProduk(newProduk)
	return row, err

}

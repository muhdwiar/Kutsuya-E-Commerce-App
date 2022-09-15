package usecase

import (
	history "project/kutsuya/features/history_order"
	queryProduk "project/kutsuya/features/produk/data"

	"gorm.io/gorm"
)

type historyUsecase struct {
	historyData history.DataInterface
}

func New(data history.DataInterface) history.UsecaseInterface {
	return &historyUsecase{
		historyData: data,
	}
}

func (usecase *historyUsecase) InsertHistoryOrder(newHistory history.CoreRequest) (int, error) {

	var db *gorm.DB
	data_produk := queryProduk.New(db)

	dataProduk, errProduk := data_produk.SelectProdukById(int(newHistory.Product_Id))

	if errProduk != nil {
		return -1, errProduk
	}

	var newHistoryCore history.Core
	newHistoryCore.User_Id = newHistory.User_Id
	newHistoryCore.Nama_Produk = dataProduk.Nama_Produk
	newHistoryCore.Jumlah = newHistory.Jumlah
	newHistoryCore.Total_Biaya = newHistory.Jumlah * dataProduk.Harga
	newHistoryCore.Status_Order = "Pembelian Sukses"

	row, err := usecase.historyData.CreateHistoryOrder(newHistoryCore)

	return row, err

}

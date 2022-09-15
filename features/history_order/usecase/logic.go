package usecase

import (
	history "project/kutsuya/features/history_order"
	cart "project/kutsuya/features/shopping_cart"
)

type historyUsecase struct {
	historyData history.DataInterface
	cartData    cart.DataInterface
}

func New(data history.DataInterface, dataCart cart.DataInterface) history.UsecaseInterface {
	return &historyUsecase{
		historyData: data,
		cartData:    dataCart,
	}
}

func (usecase *historyUsecase) InsertHistoryOrder(newHistoryList []history.CoreRequest, user_id int) (int, error) {

	var total_row, row_delete int
	for _, valueHistory := range newHistoryList {
		dataProduk, errProduk := usecase.cartData.GetCartByID(int(valueHistory.Cart_Id))
		if errProduk != nil {
			return total_row, errProduk
		}

		var newHistoryCore history.Core
		newHistoryCore.User_Id = uint(user_id)
		newHistoryCore.Nama_Produk = dataProduk.Nama_Produk
		newHistoryCore.Jumlah = valueHistory.Jumlah
		newHistoryCore.Total_Biaya = valueHistory.Jumlah * dataProduk.Harga
		newHistoryCore.Status_Order = "Pembelian Sukses"

		row, err := usecase.historyData.CreateHistoryOrder(newHistoryCore)

		total_row += row
		if err != nil {
			return total_row, err
		}

		rowDel, errDel := usecase.cartData.DelCartByID(int(valueHistory.Cart_Id), user_id)

		row_delete += rowDel
		if errDel != nil {
			return rowDel, errDel
		}

	}

	return total_row, nil

}

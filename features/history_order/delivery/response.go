package delivery

import (
	"project/kutsuya/features/history_order"
)

type HistoryResponse struct {
	ID           uint   `json:"id" form:"id"`
	User_Id      uint   `json:"user_id" form:"user_id"`
	Nama_Produk  string `json:"nama_produk" form:"nama_produk"`
	Jumlah       int    `json:"jumlah" form:"jumlah"`
	Total_Biaya  int    `json:"total_biaya" form:"total_biaya"`
	Status_Order string `json:"status_order" form:"status_order"`
}

func FromCore(dataCore history_order.Core) HistoryResponse {
	return HistoryResponse{
		ID:           dataCore.ID,
		User_Id:      dataCore.User_Id,
		Nama_Produk:  dataCore.Nama_Produk,
		Jumlah:       dataCore.Jumlah,
		Total_Biaya:  dataCore.Total_Biaya,
		Status_Order: dataCore.Status_Order,
	}

}

func FromCoreList(dataCore []history_order.Core) []HistoryResponse {
	var dataResponse []HistoryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}

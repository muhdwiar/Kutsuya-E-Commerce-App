package delivery

import (
	"project/kutsuya/features/history_order"
)

type HistoryRequest struct {
	User_Id    uint `json:"user_id" form:"user_id"`
	Product_Id uint `json:"product_id" form:"product_id"`
	Jumlah     int  `json:"jumlah" form:"jumlah"`
}

func ToCoreRequest(dataRequest HistoryRequest) history_order.CoreRequest {
	return history_order.CoreRequest{
		User_Id:    dataRequest.User_Id,
		Product_Id: dataRequest.Product_Id,
		Jumlah:     dataRequest.Jumlah,
	}

}

func ToCore(dataRequest HistoryRequest, namaProduk, status_order string, total int) history_order.Core {
	return history_order.Core{
		User_Id:      dataRequest.User_Id,
		Nama_Produk:  namaProduk,
		Jumlah:       dataRequest.Jumlah,
		Total_Biaya:  total,
		Status_Order: status_order,
	}

}

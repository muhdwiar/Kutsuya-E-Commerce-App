package delivery

import (
	"project/kutsuya/features/history_order"
)

type HistoryRequest struct {
	Cart_Id uint `json:"cart_id" form:"cart_id"`
	Jumlah  int  `json:"jumlah" form:"jumlah"`
}

func ToCoreRequest(dataRequest HistoryRequest) history_order.CoreRequest {
	return history_order.CoreRequest{
		Cart_Id: dataRequest.Cart_Id,
		Jumlah:  dataRequest.Jumlah,
	}

}

func ToCoreRequestList(dataRequest []HistoryRequest) []history_order.CoreRequest {
	var dataCore []history_order.CoreRequest
	for _, v := range dataRequest {
		dataCore = append(dataCore, ToCoreRequest(v))
	}
	return dataCore
}

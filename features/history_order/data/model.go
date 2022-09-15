package data

import (
	"project/kutsuya/features/history_order"
	userModel "project/kutsuya/features/user/data"

	"gorm.io/gorm"
)

type History_Order struct {
	gorm.Model
	User_Id      uint
	Nama_Produk  string
	Jumlah       int
	Total_Biaya  int
	Status_Order string
	User         userModel.User `gorm:"foreignKey:User_Id"`
}

func fromCore(dataCore history_order.Core) History_Order {
	return History_Order{
		User_Id:      dataCore.User_Id,
		Nama_Produk:  dataCore.Nama_Produk,
		Jumlah:       dataCore.Jumlah,
		Total_Biaya:  dataCore.Total_Biaya,
		Status_Order: dataCore.Status_Order,
	}

}

func (dataHistory *History_Order) toCore() history_order.Core {
	return history_order.Core{
		ID:           dataHistory.ID,
		User_Id:      dataHistory.User_Id,
		Nama_Produk:  dataHistory.Nama_Produk,
		Jumlah:       dataHistory.Jumlah,
		Total_Biaya:  dataHistory.Total_Biaya,
		Status_Order: dataHistory.Status_Order,
	}
}

func toCoreList(dataHistory []History_Order) []history_order.Core {
	var dataCore []history_order.Core

	for key := range dataHistory {
		dataCore = append(dataCore, dataHistory[key].toCore())

	}

	return dataCore

}

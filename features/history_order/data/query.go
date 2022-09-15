package data

import (
	history "project/kutsuya/features/history_order"

	"gorm.io/gorm"
)

type historyData struct {
	db *gorm.DB
}

func New(db *gorm.DB) history.DataInterface {
	return &historyData{
		db: db,
	}

}

func (repo *historyData) CreateHistoryOrder(newHistory history.Core) (int, error) {

	newHistoryModel := fromCore(newHistory)

	tx := repo.db.Create(&newHistoryModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *historyData) FindAllOrder(user_id int) ([]history.Core, error) {
	var all_OrderData []History_Order
	tx := repo.db.Where("user_id = ?", user_id).Find(&all_OrderData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	order_List := toCoreList(all_OrderData)
	return order_List, nil

}

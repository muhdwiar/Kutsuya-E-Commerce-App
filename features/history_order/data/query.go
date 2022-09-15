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

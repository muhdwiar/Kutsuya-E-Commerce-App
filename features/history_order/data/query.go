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

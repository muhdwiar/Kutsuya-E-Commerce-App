package usecase

import (
	history "project/kutsuya/features/history_order"
)

type historyUsecase struct {
	historyData history.DataInterface
}

func New(data history.DataInterface) history.UsecaseInterface {
	return &historyUsecase{
		historyData: data,
	}
}

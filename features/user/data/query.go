package data

import (
	"errors"
	"project/kutsuya/features/user"
	"project/kutsuya/middlewares"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}

}

func (repo *userData) InsertData(data user.Core) (string, int, error) {
	newUser := fromCore(data)
	token, errToken := middlewares.CreateToken(int(newUser.ID))
	if errToken != nil {
		return "", -1, errToken
	}

	tx := repo.db.Create(&newUser)
	if tx.Error != nil {
		return "", 0, tx.Error
	}

	return token, int(tx.RowsAffected), nil
}

func (repo *userData) LoginUser(data user.Core) (string, error) {
	var userData User
	tx := repo.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&userData)

	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return "", errToken
	}

	return token, nil

}

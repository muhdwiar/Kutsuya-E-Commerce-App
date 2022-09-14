package data

import (
	"errors"
	"project/kutsuya/features/user"
	"project/kutsuya/middlewares"

	"golang.org/x/crypto/bcrypt"
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

func (repo *userData) GetUserById(id int) (user.Core, error) {
	var userData User
	userData.ID = uint(id)
	tx := repo.db.First(&userData)

	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	return userData.toCore(), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (repo *userData) InsertData(data user.Core) (string, int, error) {
	hash_pass, errHash := HashPassword(data.Password)

	if errHash != nil {
		return "", -1, errHash
	}
	data.Password = hash_pass //memasukkan hasil enskripsi data password

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
	tx := repo.db.Where("email = ?", data.Email).First(&userData)

	check_result := CheckPasswordHash(data.Password, userData.Password)

	if !check_result {
		return "", errors.New("password salah")
	}

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

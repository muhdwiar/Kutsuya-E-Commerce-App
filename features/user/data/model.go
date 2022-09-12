package data

import (
	"project/kutsuya/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama_User string
	Email     string
	Password  string
	Alamat    string
}

func fromCore(dataCore user.Core) User {
	return User{
		Nama_User: dataCore.Nama_User,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Alamat:    dataCore.Alamat,
	}

}

func (dataUser *User) toCore() user.Core {
	return user.Core{
		ID:        dataUser.ID,
		Nama_User: dataUser.Nama_User,
		Email:     dataUser.Email,
		Password:  dataUser.Password,
		Alamat:    dataUser.Alamat,
	}
}

func toCoreList(dataUser []User) []user.Core {
	var dataCore []user.Core

	for key := range dataUser {
		dataCore = append(dataCore, dataUser[key].toCore())

	}

	return dataCore

}

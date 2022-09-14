package user

import "time"

type Core struct {
	ID         uint
	Nama_User  string
	Email      string
	Password   string
	Alamat     string
	Created_At time.Time
	Updated_At time.Time
}

type UsecaseInterface interface {
	PostData(data Core) (token string, row int, err error)
	PostLogin(data Core) (token string, err error)
	GetById(id int) (data Core, err error)
	PutUser(data Core) (row int, err error)
	DeleteUser(id int) (row int, err error)
}

type DataInterface interface {
	InsertData(data Core) (token string, row int, err error)
	LoginUser(data Core) (token string, err error)
	GetUserById(id int) (data Core, err error)
	UpdateUser(data Core) (row int, err error)
	UserDelete(id int) (row int, err error)
}

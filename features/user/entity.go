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
}

type DataInterface interface {
	InsertData(data Core) (token string, row int, err error)
	LoginUser(data Core) (token string, err error)
}

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
}

type DataInterface interface {
}

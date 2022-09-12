package delivery

import "project/kutsuya/features/user"

type UserRequest struct {
	Nama_User string `json:"nama_user" form:"nama_user"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Alamat    string `json:"alamat" form:"alamat"`
}

func ToCore(dataRequest UserRequest) user.Core {
	return user.Core{
		Nama_User: dataRequest.Nama_User,
		Email:     dataRequest.Email,
		Password:  dataRequest.Password,
		Alamat:    dataRequest.Alamat,
	}

}

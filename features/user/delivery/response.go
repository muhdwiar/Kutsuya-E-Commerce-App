package delivery

import "project/kutsuya/features/user"

type UserResponse struct {
	ID        uint   `json:"id" form:"id"`
	Nama_User string `json:"nama_user" form:"nama_user"`
	Email     string `json:"email" form:"email"`
	Alamat    string `json:"alamat" form:"alamat"`
}

func FromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:        dataCore.ID,
		Nama_User: dataCore.Nama_User,
		Email:     dataCore.Email,
		Alamat:    dataCore.Alamat,
	}

}

func FromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}

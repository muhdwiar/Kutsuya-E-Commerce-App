package usecase

import (
	"errors"
	"project/kutsuya/features/user"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) GetById(id int) (user.Core, error) {
	result, err := usecase.userData.GetUserById(id)
	if err != nil {
		return user.Core{}, err
	}
	return result, nil
}

func (usecase *userUsecase) PostData(data user.Core) (string, int, error) {
	if data.Nama_User == "" || data.Email == "" || data.Password == "" {
		return "", -1, errors.New("data input ada yang kosong")
	}

	token, row, err := usecase.userData.InsertData(data)
	if err != nil {
		return "", -1, err
	}

	return token, row, err

}

func (usecase *userUsecase) PostLogin(data user.Core) (string, error) {
	if data.Email == "" || data.Password == "" {
		return "", errors.New("data input ada yang kosong")
	}

	token, err := usecase.userData.LoginUser(data)
	if err != nil {
		return "", err
	}

	return token, err

}

func (usecase *userUsecase) PutUser(data user.Core) (int, error) {
	row, err := usecase.userData.UpdateUser(data)
	return row, err
}

func (usecase *userUsecase) DeleteUser(id int) (int, error) {
	result, err := usecase.userData.UserDelete(id)
	if err != nil {
		return -1, err
	}
	return result, nil
}

package usecase

import (
	"project/kutsuya/features/user"
	"project/kutsuya/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {
	repo := new(mocks.Core)
	dataInput := user.Core{ID: 3, Nama_User: "nama", Email: "email", Password: "password", Alamat: "alamat"}

	t.Run("Success Post Produk", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return("6ra67rx6adgiagd9gcajwd", 1, nil)
		usecase := New(repo)
		token, row, err := usecase.PostData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, token, "6ra67rx6adgiagd9gcajwd")
		assert.Equal(t, row, 1)
		repo.AssertExpectations(t)
	})
}

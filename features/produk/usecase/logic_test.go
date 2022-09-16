package usecase

import (
	"project/kutsuya/features/produk"
	"project/kutsuya/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet_AllProduk(t *testing.T) {
	repo := new(mocks.ProdukData)
	returnData := []produk.Core{{ID: 3, Nama_Produk: "Sepatu anak", Ukuran: 26, Merk: "Ardiles", Warna: "Biru", Gender_Pengguna: "Pria", Harga: 60000, Desksripsi: "sepatu keren untuk anak kecil", File_Image: "url_image.com", User_Id: 8}}

	t.Run("Success Get All Produk", func(t *testing.T) {
		repo.On("Select_AllProduk").Return(returnData, nil)
		usecase := New(repo)
		result, err := usecase.Get_AllProduk()
		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestPostProduk(t *testing.T) {
	repo := new(mocks.ProdukData)
	dataInput := produk.Core{ID: 3, Nama_Produk: "Sepatu anak", Ukuran: 26, Merk: "Ardiles", Warna: "Biru", Gender_Pengguna: "Pria", Harga: 60000, Desksripsi: "sepatu keren untuk anak kecil", File_Image: "url_image.com", User_Id: 8}

	t.Run("Success Post Produk", func(t *testing.T) {
		repo.On("InsertProduk", mock.Anything).Return(1, nil)
		usecase := New(repo)
		result, err := usecase.PostProduk(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

}

func TestGetProdukById(t *testing.T) {
	repo := new(mocks.ProdukData)
	returnData := produk.Core{ID: 1, Nama_Produk: "Sepatu anak", Ukuran: 26, Merk: "Ardiles", Warna: "Biru", Gender_Pengguna: "Pria", Harga: 60000, Desksripsi: "sepatu keren untuk anak kecil", File_Image: "url_image.com", User_Id: 8}
	product_id := 1

	t.Run("Success Get Produk By Id", func(t *testing.T) {
		repo.On("SelectProdukById", product_id).Return(returnData, nil)
		usecase := New(repo)
		result, err := usecase.GetProdukById(product_id)
		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})

}

func TestPutProduk(t *testing.T) {
	repo := new(mocks.ProdukData)
	dataInput := produk.Core{Nama_Produk: "Sepatu anak", Ukuran: 26, Merk: "Ardiles", Warna: "Biru", Gender_Pengguna: "Pria", Harga: 60000, Desksripsi: "sepatu keren untuk anak kecil", File_Image: "url_image.com", User_Id: 8}
	product_id := 1

	t.Run("Success Put Produk", func(t *testing.T) {
		repo.On("UpdateDataProduk", mock.Anything, product_id).Return(1, nil)
		usecase := New(repo)
		result, err := usecase.PutProduk(dataInput, product_id)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

}

func TestDeleteProdukUser(t *testing.T) {
	repo := new(mocks.ProdukData)
	user_id := 1
	product_id := 1

	t.Run("Success Delete Produk", func(t *testing.T) {
		repo.On("DeleteProduk", user_id, product_id).Return(1, nil)
		usecase := New(repo)
		result, err := usecase.DeleteProdukUser(user_id, product_id)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

}

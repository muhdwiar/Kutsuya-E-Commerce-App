package usecase

import (
	cart "project/kutsuya/features/shopping_cart"
	"project/kutsuya/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectCarts(t *testing.T) {
	repoCart := new(mocks.CartData)
	repoProduk := new(mocks.ProdukData)
	returnData := []cart.Core{{ID: 3, User_Id: 2, Product_Id: 1, Nama_Produk: "Sepatu anak", Ukuran: 26, Merk: "Ardiles", Harga: 60000, File_Image: "url_image.com"}}

	t.Run("Success Get Carts Data", func(t *testing.T) {
		repoCart.On("FindCarts", 2).Return(returnData, nil)
		usecase := New(repoCart, repoProduk)
		results, err := usecase.SelectCarts(2)
		assert.NoError(t, err)
		assert.Equal(t, results, returnData)
		repoCart.AssertExpectations(t)
	})
}

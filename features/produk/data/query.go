package data

import (
	"errors"
	"project/kutsuya/features/produk"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) produk.DataInterface {
	return &productData{
		db: db,
	}

}

func (repo *productData) UpdateDataProduk(dataProduk produk.Core) (int, error) {
	var dataUpdate Produk

	tx_OldData := repo.db.First(&dataUpdate, dataProduk.ID)

	if tx_OldData.Error != nil {
		return -1, tx_OldData.Error
	}

	if dataProduk.Nama_Produk != "" {
		dataUpdate.Nama_Produk = dataProduk.Nama_Produk
	}

	if dataProduk.Ukuran != 0 {
		dataUpdate.Ukuran = dataProduk.Ukuran
	}

	if dataProduk.Merk != "" {
		dataUpdate.Merk = dataProduk.Merk
	}

	if dataProduk.Warna != "" {
		dataUpdate.Warna = dataProduk.Warna
	}

	if dataProduk.Gender_Pengguna != "" {
		dataUpdate.Gender_Pengguna = dataProduk.Gender_Pengguna
	}

	if dataProduk.Harga != 0 {
		dataUpdate.Harga = dataProduk.Harga
	}

	if dataProduk.Desksripsi != "" {
		dataUpdate.Desksripsi = dataProduk.Desksripsi
	}

	tx_newData := repo.db.Save(&dataUpdate)

	if tx_newData.Error != nil {
		return -1, tx_newData.Error
	}

	if tx_newData.RowsAffected != 1 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx_newData.RowsAffected), nil

}

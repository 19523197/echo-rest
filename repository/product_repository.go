package repository

import (
	"belajar-echo/model"
	"database/sql"
)

type ProductRepoCont interface {
	GetAll() ([]model.Product, error)
	Insert(model.Product) (model.Product, error)
}

type ProductRepo struct {
	Sql *sql.DB
}

func NewProductRepo(s *sql.DB) ProductRepoCont {
	return &ProductRepo{Sql: s}
}

func (f *ProductRepo) GetAll() (produks []model.Product, err error) {
	var produk model.Product
	results, err := f.Sql.Query("SELECT nama_produk, harga_produk FROM toko")
	if err != nil {
		return nil, err
	}

	for results.Next() {
		err = results.Scan(&produk.Name, &produk.Price)
		if err != nil {
			return produks, err
		}
		produks = append(produks, produk)
	}

	return produks, err
}

func (f *ProductRepo) Insert(product model.Product) (response model.Product, err error) {
	_, err = f.Sql.Exec("INSERT INTO toko (nama_produk, harga_produk) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		return response, err
	}

	response = product
	return response, err
}

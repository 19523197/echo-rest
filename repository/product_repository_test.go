package repository

import (
	"belajar-echo/database"
	"belajar-echo/model"
	"testing"
)

func TestGetAll(t *testing.T) {
	db, err := database.SetupSQLDatabase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	repo := ProductRepo{Sql: db}
	_, err = repo.GetAll()

	if err != nil {
		t.Error(err)
	}

}

func TestInsert(t *testing.T) {
	produk := model.Product{
		Name:  "produk test",
		Price: 5000,
	}
	db, err := database.SetupSQLDatabase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	repo := ProductRepo{Sql: db}
	result, err := repo.Insert(produk)

	if result != produk {
		t.Error("Tidak sesuai")
	}
	if err != nil {
		t.Error(err)
	}

}

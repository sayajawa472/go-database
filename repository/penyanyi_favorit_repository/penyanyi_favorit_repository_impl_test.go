package penyanyi_favorit_repository

import (
	"context"
	"fmt"
	"golang_database"
	belajar_db "golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnectInsert(t *testing.T) {
	penyanyiFavoritRepository := NewPenyanyiFavoritRepository(golang_database.GetConnection())
	ctx := context.Background()
	penyanyi := entity.PenyanyiFavorit{
		Nama:          "Taylor Swift",
		Jenis_Kelamin: "wanita",
		Hobi:          "menulis lagu",
	}

	result, err := penyanyiFavoritRepository.Insert(ctx, penyanyi)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	penyanyiFavoritRepository := NewPenyanyiFavoritRepository(golang_database.GetConnection())
	penyanyiFavorit, err := penyanyiFavoritRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(penyanyiFavorit)
}

func TestFindAll(t *testing.T) {
	penyanyiFavoritRepository := NewPenyanyiFavoritRepository(golang_database.GetConnection())
	PenyanyiFavorit, err := penyanyiFavoritRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, PenyanyiFavorit := range PenyanyiFavorit {
		fmt.Println(PenyanyiFavorit)
	}
}

func TestDelete(t *testing.T) {

	penyanyiFavoritRepository := NewPenyanyiFavoritRepository(golang_database.GetConnection())
	result, err := penyanyiFavoritRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUpdate(t *testing.T) {
	penyanyiFavoritRepository := NewPenyanyiFavoritRepository(belajar_db.GetConnection())

	ctx := context.Background()
	penyanyi := entity.PenyanyiFavorit{
		Id:            1,
		Nama:          "Taylor Swift",
		Jenis_Kelamin: "wanita",
		Hobi:          "menulis lagu",
	}

	result, err := penyanyiFavoritRepository.Update(ctx, 1, penyanyi)
	if err != nil {
		panic(err)
		fmt.Println(result)
	}
}

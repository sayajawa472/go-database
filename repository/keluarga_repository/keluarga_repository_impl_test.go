package keluarga_repository

import (
	"context"
	"fmt"
	"golang_database"
	belajar_db "golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnectionInsert(t *testing.T) {
	KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())

	ctx := context.Background()
	keluarga := entity.Keluarga{
		Nama: "Khansa Zahra",
		Umur: 16,
	}

	result, err := KeluargaRepository.Insert(ctx, keluarga)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {

	KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())

	keluarga, err := KeluargaRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
		fmt.Println(keluarga)
	}
}

func TestFindAll(t *testing.T) {

	KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())

	keluarga, err := KeluargaRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
		fmt.Println(keluarga)
	}
}

func TestDelete(t *testing.T) {
	KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())
	result, err := KeluargaRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUpdate(t *testing.T) {
	KeluargaRepository := NewKeluargaRepository(belajar_db.GetConnection())

	ctx := context.Background()
	Keluarga := entity.Keluarga{
		Id:   1,
		Nama: "Khansa Zahra1",
		Umur: 17,
	}

	result, err := KeluargaRepository.Update(ctx, Keluarga)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

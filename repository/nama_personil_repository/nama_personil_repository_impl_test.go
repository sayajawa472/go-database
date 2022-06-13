package nama_personil_repository

import (
	"context"
	"fmt"
	"golang_database"
	belajar_db "golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnectInsert(t *testing.T) {
	namaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())

	ctx := context.Background()
	namaPersonil := entity.NamaPersonil{
		Nama:          "Harry Styles",
		Jenis_Kelamin: "Pria",
	}

	result, err := namaPersonilRepository.Insert(ctx, namaPersonil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())
	namapersonil, err := NamaPersonilRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(namapersonil)
}

func TestFindAll(t *testing.T) {
	NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())
	NamaPersonil, err := NamaPersonilRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, NamaPersonil := range NamaPersonil {
		fmt.Println(NamaPersonil)
	}
}

func TestDelete(t *testing.T) {

	NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())
	result, err := NamaPersonilRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUpdate(t *testing.T) {
	NamaPersonilRepository := NewNamaPersonilRepository(belajar_db.GetConnection())

	ctx := context.Background()
	NamaPersonil := entity.NamaPersonil{
		Id:            1,
		Nama:          "Harry Styles",
		Jenis_Kelamin: "Pria",
	}

	result, err := NamaPersonilRepository.Update(ctx, 1, NamaPersonil)
	if err != nil {
		panic(err)
		fmt.Println(result)
	}
}

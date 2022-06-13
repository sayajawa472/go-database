package perlengkapan_sekolah_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnextionInsert(t *testing.T) {
	perlengkapanSekolahRepository := NewPerlengkapanSekolahRepository(golang_database.GetConnection())

	ctx := context.Background()
	perlengkapanSekolah := entity.PerlengkapanSekolah{
		Nama:   "ransel",
		Bahan:  "Kanvas",
		Fungsi: "untuk membawa barang barang",
		Harga:  "100",
	}

	result, err := perlengkapanSekolahRepository.Insert(ctx, perlengkapanSekolah)
	if err != nil {

		fmt.Println(result)
	}
}

func TestFindById(t *testing.T) {

	perlengkapanSekolahRepository := NewPerlengkapanSekolahRepository(golang_database.GetConnection())

	perlengkapanSekolah, err := perlengkapanSekolahRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
		fmt.Println(perlengkapanSekolah)
	}
}

func TestFindAll(t *testing.T) {

	perlengkapanSekolahRepository := NewPerlengkapanSekolahRepository(golang_database.GetConnection())

	perlengkapanSekolah, err := perlengkapanSekolahRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
		fmt.Println(perlengkapanSekolah)
	}
}

func TestDelete(t *testing.T) {
	perlengkapanSekolahRepository := NewPerlengkapanSekolahRepository(golang_database.GetConnection())
	result, err := perlengkapanSekolahRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)

	}
	fmt.Println(result)
}

func TestUpdate(t *testing.T) {
	perlengkapanSekolahRepository := NewPerlengkapanSekolahRepository(golang_database.GetConnection())

	perlengkapanSekolah, err := perlengkapanSekolahRepository.Delete(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(perlengkapanSekolah)
}

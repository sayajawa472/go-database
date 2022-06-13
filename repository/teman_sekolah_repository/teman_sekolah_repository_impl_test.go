package teman_sekolah_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnextionInsert(t *testing.T) {
	temanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

	ctx := context.Background()
	temanSekolah := entity.TemanSekolah{
		Nama:   "irmawati",
		Alamat: "kapuk",
		Hobi:   "berenang",
	}

	result, err := temanSekolahRepository.Insert(ctx, temanSekolah)
	if err != nil {

		fmt.Println(result)
	}
}

func TestFindById(t *testing.T) {

	temanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

	temanSekolah, err := temanSekolahRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
		fmt.Println(temanSekolah)
	}
}

func TestFindAll(t *testing.T) {

	temanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

	temanSekolah, err := temanSekolahRepository.FindByAll(context.Background())
	if err != nil {
		panic(err)
		fmt.Println(temanSekolah)
	}
}

func TestDelete(t *testing.T) {
	temanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())
	result, err := temanSekolahRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)

	}
	fmt.Println(result)
}

func TestUpdate(t *testing.T) {
	temanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

	temanSekolah, err := temanSekolahRepository.Delete(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(temanSekolah)
}

package teman_sekolah_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type TemanSekolahRepositoryImpl struct {
	DB *sql.DB
}

func NewTemanSekolahRepository(db *sql.DB) TemanSekolah {
	return &TemanSekolahRepositoryImpl{DB: db}
}

func (repository *TemanSekolahRepositoryImpl) Insert(ctx context.Context, TemanSekolah entity.TemanSekolah) (entity.TemanSekolah, error) {
	script := "INSERT INTO teman_sekolah(nama, alamat, hobi) VALUE(?, ?,?)"
	result, err := repository.DB.ExecContext(ctx, script, TemanSekolah.Nama, TemanSekolah.Alamat, TemanSekolah.Hobi)
	if err != nil {
		return TemanSekolah, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return TemanSekolah, err
	}
	TemanSekolah.Id = int32(id)
	return TemanSekolah, nil
}

func (repository *TemanSekolahRepositoryImpl) FindById(ctx context.Context, Id int32) (entity.TemanSekolah, error) {
	script := "SELECT id, nama, alamat, hobi FROM teman_sekolah WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, Id)
	temanSekolah := entity.TemanSekolah{}

	if err != nil {
		return temanSekolah, err
	}
	defer rows.Close()
	if rows.Next() {
		// ADA
		rows.Scan(&temanSekolah.Id, &temanSekolah.Nama, &temanSekolah.Alamat, &temanSekolah.Hobi)
		return temanSekolah, nil
	} else {
		//tidak ada
		return temanSekolah, errors.New("id " + strconv.Itoa(int(Id)) + "not found")
	}
}

func (repository *TemanSekolahRepositoryImpl) FindByAll(ctx context.Context) ([]entity.TemanSekolah, error) {
	script := "SELECT id, nama, alamat, hobi FROM teman_sekolah"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var temanSekolah []entity.TemanSekolah
	for rows.Next() {
		teman := entity.TemanSekolah{}
		rows.Scan(&teman.Id, &teman.Nama, &teman.Alamat, &teman.Hobi)
		temanSekolah = append(temanSekolah, teman)
	}
	return temanSekolah, nil
}

func (repository *TemanSekolahRepositoryImpl) Update(ctx context.Context, TemanSekolah entity.TemanSekolah) (entity.TemanSekolah, error) {
	script := "UPDATE teman_sekolah SET nama= ?, alamat= ?, hobi= ? WHERE id= ?"
	result, err := repository.DB.ExecContext(ctx, script, TemanSekolah.Nama, TemanSekolah.Alamat, TemanSekolah.Hobi)
	if err != nil {
		return TemanSekolah, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return TemanSekolah, err
	}
	TemanSekolah.Id = int32(id)
	return TemanSekolah, nil
}

func (repository *TemanSekolahRepositoryImpl) Delete(ctx context.Context, Id int32) (entity.TemanSekolah, error) {
	// string query ke database
	script := "DELETE teman_sekolah WHERE id = ? LIMIT 1"
	//panggil method dari QueryContext dari interface DB, yg mengembalikan
	// *rows, dan Error, *Rows => row, error => err
	_, err := repository.DB.ExecContext(ctx, script, Id)
	// new object perlengkapanSekolah baru, atau objek kosong
	temanSekolah := entity.TemanSekolah{}
	//check, apakah variable err kosong?
	if err != nil {
		//kalau variable err tidak kosong
		return temanSekolah, err
	}

	//jika berhasil dan tidak ada error, maka update colom commentId dengan id yg terakhir
	temanSekolah.Id = Id
	//kembalikan nilai perlengkapanSekolah dan error nya nill atau kosong
	return temanSekolah, nil
}

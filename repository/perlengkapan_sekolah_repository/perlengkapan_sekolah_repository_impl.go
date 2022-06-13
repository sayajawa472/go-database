package perlengkapan_sekolah_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type PerlengkapanSekolahImpl struct {
	DB *sql.DB
}

func NewPerlengkapanSekolahRepository(db *sql.DB) PerlengkapanSekolahRepository {
	return &PerlengkapanSekolahImpl{DB: db}
}

func (repository *PerlengkapanSekolahImpl) Insert(ctx context.Context, perlengkapanSekolah entity.PerlengkapanSekolah) (entity.PerlengkapanSekolah, error) {
	script := "INSERT INTO perlengkapan_sekolah(nama, bahan, fungsi, harga) VALUE(?, ?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, perlengkapanSekolah.Nama, perlengkapanSekolah.Bahan, perlengkapanSekolah.Fungsi, perlengkapanSekolah.Harga)
	if err != nil {
		return perlengkapanSekolah, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return perlengkapanSekolah, err
	}
	perlengkapanSekolah.Id = int32(id)
	return perlengkapanSekolah, nil
}

func (repository *PerlengkapanSekolahImpl) FindById(ctx context.Context, id int32) (entity.PerlengkapanSekolah, error) {
	script := "SELECT id, nama, bahan, fungsi, harga, FROM perlengkpan_sekolah WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	perlengkapanSekolah := entity.PerlengkapanSekolah{}

	if err != nil {
		return perlengkapanSekolah, err
	}
	defer rows.Close()
	if rows.Next() {
		// ADA
		rows.Scan(&perlengkapanSekolah.Id, &perlengkapanSekolah.Nama, &perlengkapanSekolah.Bahan, &perlengkapanSekolah.Fungsi, &perlengkapanSekolah.Harga)
		return perlengkapanSekolah, nil
	} else {
		//tidak ada
		return perlengkapanSekolah, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *PerlengkapanSekolahImpl) FindAll(ctx context.Context) ([]entity.PerlengkapanSekolah, error) {
	script := "SELECT id, nama, bahan, fungsi, harga FROM perlengkapan_sekolah"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var perlengkapanSekolah []entity.PerlengkapanSekolah
	for rows.Next() {
		perlengkapan := entity.PerlengkapanSekolah{}
		rows.Scan(&perlengkapan.Id, &perlengkapan.Nama, &perlengkapan.Bahan, &perlengkapan.Fungsi, &perlengkapan.Harga)
		perlengkapanSekolah = append(perlengkapanSekolah, perlengkapan)
	}
	return perlengkapanSekolah, nil
}

func (repository *PerlengkapanSekolahImpl) Update(ctx context.Context, perlengkapanSekolah entity.PerlengkapanSekolah) (entity.PerlengkapanSekolah, error) {
	script := "UPDATE perlengkapan_sekolah SET nama= ?, bahan= ?, fungsi= ?, harga= ? WHERE id= ?"
	result, err := repository.DB.ExecContext(ctx, script, perlengkapanSekolah.Nama, perlengkapanSekolah.Bahan, perlengkapanSekolah.Fungsi, perlengkapanSekolah.Harga)
	if err != nil {
		return perlengkapanSekolah, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return perlengkapanSekolah, err
	}
	perlengkapanSekolah.Id = int32(id)
	return perlengkapanSekolah, nil
}

func (repository *PerlengkapanSekolahImpl) Delete(ctx context.Context, id int32) (entity.PerlengkapanSekolah, error) {
	// string query ke database
	script := "DELETE perlengkapan_sekolah WHERE id = ? LIMIT 1"
	//panggil method dari QueryContext dari interface DB, yg mengembalikan
	// *rows, dan Error, *Rows => row, error => err
	_, err := repository.DB.ExecContext(ctx, script, id)
	// new object perlengkapanSekolah baru, atau objek kosong
	perlengkapanSekolah := entity.PerlengkapanSekolah{}
	//check, apakah variable err kosong?
	if err != nil {
		//kalau variable err tidak kosong
		return perlengkapanSekolah, err
	}

	//jika berhasil dan tidak ada error, maka update colom commentId dengan id yg terakhir
	perlengkapanSekolah.Id = id
	//kembalikan nilai perlengkapanSekolah dan error nya nill atau kosong
	return perlengkapanSekolah, nil
}

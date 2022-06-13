package nama_personil_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type NamaPersonilRepositoryImpl struct {
	DB *sql.DB
}

func NewNamaPersonilRepository(db *sql.DB) *NamaPersonilRepositoryImpl {
	return &NamaPersonilRepositoryImpl{DB: db}
}

func (repository *NamaPersonilRepositoryImpl) Insert(ctx context.Context, namaPersonil entity.NamaPersonil) (entity.NamaPersonil, error) {
	script := "INSERT INTO nama_personil(nama, jenis_kelamin) VALUE (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, namaPersonil.Nama, namaPersonil.Jenis_Kelamin)
	if err != nil {
		return namaPersonil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return namaPersonil, err

	}
	namaPersonil.Id = int32(id)
	return namaPersonil, nil
}

func (repository *NamaPersonilRepositoryImpl) FindById(ctx context.Context, id int32) (entity.NamaPersonil, error) {
	script := "SELECT id, nama, JenisKelamin FROM nama_personil WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	namaPersonil := entity.NamaPersonil{}

	if err != nil {
		return namaPersonil, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&namaPersonil.Id, &namaPersonil.Nama, &namaPersonil.Jenis_Kelamin)
		return namaPersonil, nil
	} else {
		return namaPersonil, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *NamaPersonilRepositoryImpl) FindAll(ctx context.Context) ([]entity.NamaPersonil, error) {
	script := "SELECT id, nama, JenisKelamin FROM NamaPersonil"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var namaPersonil []entity.NamaPersonil
	for rows.Next() {
		nama := entity.NamaPersonil{}
		rows.Scan(&nama.Id, &nama.Nama, &nama.Jenis_Kelamin)
		namaPersonil = append(namaPersonil, nama)
	}
	return namaPersonil, nil
}

func (repository *NamaPersonilRepositoryImpl) Update(ctx context.Context, id int32, namaPersonil entity.NamaPersonil) (entity.NamaPersonil, error) {
	script := "UPDATE NamaPersonil SET id = ?, Nama = ?, JenisKelamin = ? WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return namaPersonil, err
	}

	if rows.Next() {
		script := "UPDATE NamaPersonil SET Nama = ?, JenisKelamin = ? WHERE ID = ?"
		_, err := repository.DB.ExecContext(ctx, script, namaPersonil.Nama, namaPersonil.Jenis_Kelamin, id)
		if err != nil {
			return namaPersonil, err
		}
		namaPersonil.Id = id
		return namaPersonil, nil
	} else {
		return namaPersonil, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *NamaPersonilRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM NamaPersonil WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return id, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return id, err
	}
	if rowCnt == 0 {
		return id, err
	}
	return id, nil
}

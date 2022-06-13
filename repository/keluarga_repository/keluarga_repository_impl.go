package keluarga_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type KeluargaRepositoryImpl struct {
	DB *sql.DB
}

func NewKeluargaRepository(db *sql.DB) KeluargaRepository {
	return &KeluargaRepositoryImpl{DB: db}
}

func (repository *KeluargaRepositoryImpl) Insert(ctx context.Context, keluarga entity.Keluarga) (entity.Keluarga, error) {
	script := "INSERT INTO keluarga(nama, umur) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, keluarga.Umur, keluarga.Nama)
	if err != nil {
		return keluarga, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return keluarga, err
	}
	keluarga.Id = int32(id)
	return keluarga, nil
}

func (repository *KeluargaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Keluarga, error) {
	script := "SELECT id, nama, umur FROM keluarga"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Keluargi []entity.Keluarga
	for rows.Next() {
		Keluarga := entity.Keluarga{}
		rows.Scan(&Keluarga.Id, &Keluarga.Nama, &Keluarga.Umur)
		Keluargi = append(Keluargi, Keluarga)
	}
	return Keluargi, nil
}

func (repository *KeluargaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Keluarga, error) {
	script := "SELECT id, nama, umur, hobi FROM keluarga WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	keluarga := entity.Keluarga{}

	if err != nil {
		return keluarga, err
	}
	defer rows.Close()
	if rows.Next() {
		// ADA
		rows.Scan(&keluarga.Id, &keluarga.Nama, &keluarga.Umur, &keluarga.Hobi)
		return keluarga, nil
	} else {
		//tidak ada
		return keluarga, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *KeluargaRepositoryImpl) Update(ctx context.Context, keluarga entity.Keluarga) (entity.Keluarga, error) {
	script := "UPDATE keluarga SET nama = ?, umur = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, keluarga.Nama, keluarga.Umur, keluarga.Id)
	if err != nil {
		return keluarga, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return keluarga, err
	}
	if rowCnt == 0 {
		return keluarga, err
	}
	return keluarga, err
}

func (repository *KeluargaRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM keluarga WHERE id = ?"
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

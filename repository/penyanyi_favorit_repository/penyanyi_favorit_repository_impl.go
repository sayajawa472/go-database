package penyanyi_favorit_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type PenyanyiFavoritRepositoryImpl struct {
	DB *sql.DB
}

func NewPenyanyiFavoritRepository(db *sql.DB) *PenyanyiFavoritRepositoryImpl {
	return &PenyanyiFavoritRepositoryImpl{DB: db}
}

func (repository *PenyanyiFavoritRepositoryImpl) Insert(ctx context.Context, penyanyiFavorit entity.PenyanyiFavorit) (entity.PenyanyiFavorit, error) {
	script := "INSERT INTO penyanyi_favorit(nama, jenis_kelamin) VALUE (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, penyanyiFavorit.Nama, penyanyiFavorit.Jenis_Kelamin)
	if err != nil {
		return penyanyiFavorit, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return penyanyiFavorit, err

	}
	penyanyiFavorit.Id = int32(id)
	return penyanyiFavorit, nil
}

func (repository *PenyanyiFavoritRepositoryImpl) FindById(ctx context.Context, id int32) (entity.PenyanyiFavorit, error) {
	script := "SELECT id, nama, jeniskelamin FROM penyanyi_favorit WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	penyanyiFavorit := entity.PenyanyiFavorit{}

	if err != nil {
		return penyanyiFavorit, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&penyanyiFavorit.Id, &penyanyiFavorit.Nama, &penyanyiFavorit.Jenis_Kelamin)
		return penyanyiFavorit, nil
	} else {
		return penyanyiFavorit, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *PenyanyiFavoritRepositoryImpl) FindAll(ctx context.Context) ([]entity.PenyanyiFavorit, error) {
	script := "SELECT id, nama, jenis_kelamin FROM penyanyi_favorit"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var penyanyiFavorit []entity.PenyanyiFavorit
	for rows.Next() {
		nama := entity.PenyanyiFavorit{}
		rows.Scan(&nama.Id, &nama.Nama, &nama.Jenis_Kelamin)
		penyanyiFavorit = append(penyanyiFavorit, nama)
	}
	return penyanyiFavorit, nil
}

func (repository *PenyanyiFavoritRepositoryImpl) Update(ctx context.Context, id int32, penyanyiFavorit entity.PenyanyiFavorit) (entity.PenyanyiFavorit, error) {
	script := "UPDATE penyanyi_favorit SET id = ?, nama = ?, jenis_kelamin = ? WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return penyanyiFavorit, err
	}

	if rows.Next() {
		script := "UPDATE penyanyi_favorit SET nama = ?, jenis_kelamin = ? WHERE ID = ?"
		_, err := repository.DB.ExecContext(ctx, script, penyanyiFavorit.Nama, penyanyiFavorit.Jenis_Kelamin, id)
		if err != nil {
			return penyanyiFavorit, err
		}
		penyanyiFavorit.Id = id
		return penyanyiFavorit, nil
	} else {
		return penyanyiFavorit, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *PenyanyiFavoritRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM penyanyi_favorit WHERE id = ?"
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

package nama_personil_repository

import (
	"context"
	"golang_database/entity"
)

type NamaPersonilRepository interface {
	Insert(ctx context.Context, namaPersonil entity.NamaPersonil) (entity.NamaPersonil, error)
	FindById(ctx context.Context, id int32) (entity.NamaPersonil, error)
	FindAll(ctx context.Context) ([]entity.NamaPersonil, error)
	Update(ctx context.Context, id int32, namaPersonil entity.NamaPersonil) (entity.NamaPersonil, error)
	Delete(ctx context.Context, namaPersonil entity.NamaPersonil) (entity.NamaPersonil, error)
}

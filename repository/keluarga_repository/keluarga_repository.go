package keluarga_repository

import (
	"context"
	"golang_database/entity"
)

type KeluargaRepository interface {
	Insert(ctx context.Context, keluarga entity.Keluarga) (entity.Keluarga, error)
	FindById(ctx context.Context, Id int32) (entity.Keluarga, error)
	FindAll(ctx context.Context) ([]entity.Keluarga, error)
	Update(ctx context.Context, keluarga entity.Keluarga) (entity.Keluarga, error)
	Delete(ctx context.Context, id int32) (int32, error)
}

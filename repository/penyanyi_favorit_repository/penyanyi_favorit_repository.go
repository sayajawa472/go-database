package penyanyi_favorit_repository

import (
	"context"
	"golang_database/entity"
)

type PenyanyiFavoritRepository interface {
	Insert(ctx context.Context, penyanyiFavorit entity.PenyanyiFavorit) (entity.PenyanyiFavorit, error)
	FindById(ctx context.Context, id int32) (entity.PenyanyiFavorit, error)
	FindAll(ctx context.Context) ([]entity.PenyanyiFavorit, error)
	Update(ctx context.Context, id int32, penyanyiFavorit entity.PenyanyiFavorit) (entity.PenyanyiFavorit, error)
	Delete(ctx context.Context, penyanyiFavorit entity.PenyanyiFavorit) (entity.PenyanyiFavorit, error)
}

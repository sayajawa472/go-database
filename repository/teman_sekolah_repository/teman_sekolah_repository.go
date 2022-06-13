package teman_sekolah_repository

import (
	"context"
	"golang_database/entity"
)

type TemanSekolah interface {
	Insert(ctx context.Context, TemanSekolah entity.TemanSekolah) (entity.TemanSekolah, error)
	FindById(ctx context.Context, Id int32) (entity.TemanSekolah, error)
	FindByAll(ctx context.Context) ([]entity.TemanSekolah, error)
	Update(ctx context.Context, TemanSekolah entity.TemanSekolah) (entity.TemanSekolah, error)
	Delete(ctx context.Context, Id int32) (entity.TemanSekolah, error)
}

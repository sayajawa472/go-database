package perlengkapan_sekolah_repository

import (
	"context"
	"golang_database/entity"
)

type PerlengkapanSekolahRepository interface {
	Insert(ctx context.Context, perlengkapanSekolah entity.PerlengkapanSekolah) (entity.PerlengkapanSekolah, error)
	FindById(ctx context.Context, Id int32) (entity.PerlengkapanSekolah, error)
	FindAll(ctx context.Context) ([]entity.PerlengkapanSekolah, error)
	Update(ctx context.Context, perlengkapanSekolah entity.PerlengkapanSekolah) (entity.PerlengkapanSekolah, error)
	Delete(ctx context.Context, id int32) (entity.PerlengkapanSekolah, error)
}

package comment_repository

import (
	"context"
	"golang_database/entity"
)

type Repository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	Update(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Delete(ctx context.Context, id int32) (int32, error)
}

package comment_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnectInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	comment, err := CommentRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAllBarang(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestConnectUpdate(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := CommentRepository.Update(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	comment, err := CommentRepository.Delete(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

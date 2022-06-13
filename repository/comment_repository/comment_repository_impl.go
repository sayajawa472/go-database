package comment_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepositoryImpl {
	return CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comment(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil

}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ADA
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//tidak ada
		return comment, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *CommentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "UPDATE comments SET email= ?, comment= ? WHERE Id= ?"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Id, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, id int32) (entity.Comment, error) {
	// string query ke database
	script := "DELETE comments WHERE id = ? LIMIT 1"
	//panggil method dari QueryContext dari interface DB, yg mengembalikan
	// *rows, dan Error, *Rows => row, error => err
	_, err := repository.DB.ExecContext(ctx, script, id)
	// new object comment baru, atau objek kosong
	comment := entity.Comment{}
	//check, apakah variable err kosong?
	if err != nil {
		//kalau variable err tidak kosong
		return comment, err
	}

	//jika berhasil dan tidak ada error, maka update colom commentId dengan id yg terakhir
	comment.Id = id
	//kembalikan nilai comment dan error nya nill atau kosong
	return comment, nil
}

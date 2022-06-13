package sql

import (
	"context"
	"fmt"
	"golang_database"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "aqilla@gmail.com"
	comment := "Test komen"

	script := "INSERT INTO comments(email, comment) VALUE(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new comment with id", insertId)
}

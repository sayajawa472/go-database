package sql

import (
	"context"
	"fmt"
	"golang_database"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments (email, comment) VALUE(?, ?)"
	//do transaction
	for i := 0; i < 10; i++ {
		email := "aqilla" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}

package sql

import (
	"context"
	"fmt"
	"golang_database"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comment(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "aqilla" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}
}

package sql

import (
	"context"
	"fmt"
	"golang_database"
	"testing"
)

func TestExecSqlParameter(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "nadiyah'; DROP TABLE user; #"
	password := "nadiyah"

	script := "INSERT INTO user(username, password) VALUE(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new user")
}

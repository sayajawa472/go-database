package sql_test

import (
	"context"
	"fmt"
	"golang_database"
	"testing"
)

func TestExectSql(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUE('harry', 'Harry')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql1(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customers"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

package sql

import (
	"context"
	"fmt"
	"golang_database"
	"testing"
)

func TestQuerySql(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELEC id, name FROM cutomer"
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
		fmt.Println("id:", id)
		fmt.Println("Name:", name)
	}
}

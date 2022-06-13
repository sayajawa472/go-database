package sql

import (
	"context"
	"database/sql"
	"fmt"
	"golang_database"
	"testing"
	"time"
)

func TestQuerySqlComplex(t *testing.T) {
	db := golang_database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		defer rows.Close()

		for rows.Next() {
			var id, name string
			var email sql.NullString
			var balance int32
			var rating float64
			var birthDate sql.NullTime
			var createdAt time.Time
			var married bool

			err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
			if err != nil {
				panic(err)
			}
			fmt.Println("===============")
			fmt.Println("id:", id)
			fmt.Println("Name:", name)
			if email.Valid {
				fmt.Println("Email:", email.String)
			}
			fmt.Println("Balance:", balance)
			fmt.Println("Rating:", rating)
			if birthDate.Valid {
				fmt.Println("Email:", email.String)
			}
			fmt.Println("Married:", married)
			fmt.Println("Created At:", createdAt)
		}
	}
}

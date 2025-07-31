package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES ('nathan', 'Nathan')")

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string        // CHAR, VARCHAR
		var email sql.NullString   // Nullable String
		var balance int32          // INT
		var rating float64         // DOUBLE
		var birthDate sql.NullTime // Nullable Time
		var createdAt time.Time    // DATE, TIMESTAMP
		var married bool           // BOOLEAN
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("===================================================")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
		fmt.Println("===================================================")
		fmt.Println("")

		/*
				ID: garzya
			Name: Garzya
			Email: garzya@gmail.com
			Balance: 1000000
			Rating: 90
			Birth Date: 1999-09-09 00:00:00 +0000 UTC
			Married: true
			Created At: 2025-07-31 18:25:09 +0000 UTC
			===================================================

			===================================================
			ID: nathan
			Name: Nathan
			Email: nathan@gmail.com
			Balance: 500000
			Rating: 88.5
			Birth Date: 1996-06-12 00:00:00 +0000 UTC
			Married: true
			Created At: 2025-07-31 18:25:09 +0000 UTC
			===================================================

			===================================================
			ID: santoso
			Name: Santoso
			Balance: 759000
			Rating: 85.3
			Married: false
			Created At: 2025-07-31 18:27:22 +0000 UTC
			===================================================
		*/
	}
}

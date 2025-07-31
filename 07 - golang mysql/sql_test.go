package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var _username string
		err := rows.Scan(&_username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Username:", _username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	// Kode: Query dengan Parameter
	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var _username string
		err := rows.Scan(&_username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Username:", _username)
	} else {
		fmt.Println("Login Failed")
	}
	/*
		Login Failed
		--- PASS: TestSqlInjectionSafe (0.00s)
	*/
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "nathan"
	password := "nathan"

	ctx := context.Background()
	script := "INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "nathan@gmail.com"
	comment := "test comment lagi"

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer with ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// Define Prepare Statement
	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	// Execute the statement
	for i := 0; i < 10; i++ {
		email := "nathan" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID:", insertId)
	}
}

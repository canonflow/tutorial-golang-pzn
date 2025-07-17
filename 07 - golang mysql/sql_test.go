package main

import (
	"context"
	"fmt"
	"testing"
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

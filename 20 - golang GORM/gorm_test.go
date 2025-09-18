package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db.Exec("TRUNCATE TABLE users")
	user := User{
		ID:       "1",
		Password: "rahasia",
		Name: Name{
			FirstName:  "Nathan",
			MiddleName: "Garzya",
			LastName:   "Santoso",
		},
		Information: "Ini akan di-ignore oleh GORM",
	}

	response := db.Create(&user)

	assert.Nil(t, response.Error)
	assert.Equal(t, 1, int(response.RowsAffected))

	/*
		=== RUN   TestCreateUser
		--- PASS: TestCreateUser (0.01s)
		PASS
	*/
}

func TestBatchInsert(t *testing.T) {
	db.Exec("TRUNCATE TABLE users")

	var users []User

	for i := 1; i <= 10; i++ {
		users = append(users, User{
			ID: strconv.Itoa(i),
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
			Password: "rahasia",
		})
	}

	result := db.Create(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 10, int(result.RowsAffected))

	/*
		=== RUN   TestBatchInsert
		--- PASS: TestBatchInsert (0.01s)
		PASS
	*/
}

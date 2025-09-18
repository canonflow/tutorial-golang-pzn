package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "User 12"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)

	/*
		=== RUN   TestTransaction
		--- PASS: TestTransaction (0.01s)
		PASS
	*/
}

func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "User 12"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)

	/*
		=== RUN   TestTransactionError
		--- PASS: TestTransactionError (0.00s)
		PASS
	*/
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error

	assert.Nil(t, err)

	err = tx.Create(&User{ID: "15", Password: "rahasia", Name: Name{FirstName: "User 15"}}).Error

	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}

	/*
		=== RUN   TestManualTransactionSuccess
		--- PASS: TestManualTransactionSuccess (0.00s)
		PASS
	*/
}

func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "16", Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error

	assert.Nil(t, err)

	err = tx.Create(&User{ID: "15", Password: "rahasia", Name: Name{FirstName: "User 15"}}).Error

	assert.NotNil(t, err)

	if err == nil {
		tx.Commit()
	}

	/*
		=== RUN   TestManualTransactionError
		--- PASS: TestManualTransactionError (0.00s)
		PASS
	*/
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	result := db.First(&user)

	assert.Nil(t, result.Error)
	assert.Equal(t, "1", user.ID)

	user = User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "9", user.ID)

	/*
		=== RUN   TestQuerySingleObject
		--- PASS: TestQuerySingleObject (0.00s)
		PASS
	*/
}

func TestQueryInlineCondition(t *testing.T) {
	user := User{}

	result := db.First(&user, "id = ?", "5")
	assert.Nil(t, result.Error)
	assert.Equal(t, "5", user.ID)

	/*
		=== RUN   TestQueryInlineCondition
		--- PASS: TestQueryInlineCondition (0.00s)
		PASS
	*/
}

func TestQueryAllObjects(t *testing.T) {
	var users []User
	result := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"})

	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))

	/*
		=== RUN   TestQueryAllObjects
		--- PASS: TestQueryAllObjects (0.00s)
		PASS
	*/
}

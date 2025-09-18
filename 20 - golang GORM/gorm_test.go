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

func TestQueryCondition(t *testing.T) {
	var users []User

	result := db.Where("first_name LIKE ?", "%User%").
		Where("password = ?", "rahasia").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 15, len(users))

	/*
		=== RUN   TestQueryCondition
		--- PASS: TestQueryCondition (0.00s)
		PASS
	*/
}

func TestOrCondition(t *testing.T) {
	var users []User

	result := db.Where("first_name LIKE ?", "%User%").
		Or("password = ?", "rahasia").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 15, len(users))

	/*
		=== RUN   TestOrCondition (0.00s)
		--- PASS: TestOrCondition (0.00s)
		PASS
	*/
}

func TestNotCondition(t *testing.T) {
	var users []User

	result := db.Not("first_name LIKE ?", "%User%").
		Where("password = ?", "rahasia").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(users))

	/*
		=== RUN   TestNotCondition (0.00s)
		--- PASS: TestNotCondition (0.00s)
		PASS
	*/
}

func TestSelectFields(t *testing.T) {
	var users []User

	result := db.Select("id", "first_name").Find(&users)
	assert.Nil(t, result.Error)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.Equal(t, "", user.Name.MiddleName)
	}

	assert.Equal(t, 15, len(users))

	/*
		=== RUN   TestSelectFields (0.00s)
		--- PASS: TestSelectFields (0.00s)
		PASS
	*/
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
		},
	}

	var users []User

	result := db.Where(userCondition).Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))

	/*
		=== RUN   TestStructCondition (0.08s)
		--- PASS: TestStructCondition (0.08s)
		PASS
	*/
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
	}

	var users []User
	result := db.Where(mapCondition).Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 15, len(users))

	/*
		=== RUN   TestMapCondition
		--- PASS: TestMapCondition (0.00s)
		PASS
	*/
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User

	result := db.Order("id asc, first_name asc").
		Limit(5).
		Offset(5).
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, "14", users[0].ID)

	/*
		=== RUN   TestOrderLimitOffset
		--- PASS: TestOrderLimitOffset (0.00s)
		PASS
	*/
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse

	result := db.Model(&User{}).Select("id", "first_name", "last_name").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 15, len(users))

	/*
		=== RUN   TestQueryNonModel
		--- PASS: TestQueryNonModel (0.00s)
		PASS
	*/
}

func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)

	user.Name.FirstName = "Nathan"
	user.Name.MiddleName = "Garzya"
	user.Name.LastName = "Santoso"
	user.Password = "password123"

	result = db.Save(&user)

	/*
		=== RUN TestUpdate
		--- PASS: TestUpdate (0.00s)
		PASS
	*/
}

func TestSelectedColumns(t *testing.T) {
	// Update selected columns only with map
	result := db.Model(&User{}).Where("id = ?", "1").
		Updates(map[string]interface{}{
			"middle_name": "",
			"last_name":   "Morro",
		})

	assert.Nil(t, result.Error)

	// Update selected column
	result = db.Model(&User{}).Where("id = ?", "1").
		Update("password", "ubahlagi")

	assert.Nil(t, result.Error)

	// Update selected columns only with struct (model)
	result = db.Where("id = ?", "1").
		Updates(User{
			Name: Name{
				FirstName: "Nathan",
				LastName:  "Santoso",
			},
		})

	assert.Nil(t, result.Error)

	/*
		--- PASS: TestSelectedColumns (0.00s)
		=== RUN   TestSelectedColumns
		PASS
	*/
}

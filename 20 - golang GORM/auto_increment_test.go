package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

type UserLog struct {
	ID        int    `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *UserLog) TableName() string {
	return "user_logs"
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserID: "1",
			Action: "Test Action",
		}

		result := db.Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}

	/*
		=== RUN   TestAutoIncrement
		--- PASS: TestAutoIncrement (0.01s)
		PASS
	*/
}

func TestSaveOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserID: "1",
		Action: "Test Action",
	}

	result := db.Save(&userLog) // create
	assert.Nil(t, result.Error)

	userLog.UserID = "2"
	result = db.Save(&userLog) // update
	assert.Nil(t, result.Error)
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		ID: "99",
		Name: Name{
			FirstName: "User 99",
		},
	}

	result := db.Save(&user) // create
	assert.Nil(t, result.Error)

	user.Name.FirstName = "User 99 Updated"
	result = db.Save(&user) // update
	assert.Nil(t, result.Error)
}

func TestConflict(t *testing.T) {
	user := User{
		ID: "88",
		Name: Name{
			FirstName: "User 88 Updated",
		},
	}

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Save(&user) // create

	assert.Nil(t, result.Error)
	/*
		=== RUN   TestConflict
		--- PASS: TestConflict (0.00s)
		PASS
	*/
}

func TestDelete(t *testing.T) {
	var user User
	// Pakai model yang sudah ada ID-nya
	result := db.First(&user, "id = ?", "88")
	assert.Nil(t, result.Error)

	result = db.Delete(&user)
	assert.Nil(t, result.Error)

	// Pakai model, ambil id -> langsung delete (pakai inline where)
	result = db.Delete(&User{}, "id = ?", "99")
	assert.Nil(t, result.Error)

	// Pakai where method -> lalu pass strut User (untuk memberi tahu table apa yang dipakai)
	result = db.Where("id = ?", "15").Delete(&User{})
	assert.Nil(t, result.Error)

	/*
		=== RUN   TestDelete
		--- PASS: TestDelete (0.00s)
		PASS
	*/
}

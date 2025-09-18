package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type UserLog struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    string    `gorm:"column:user_id"`
	Action    string    `gorm:"column:action"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
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

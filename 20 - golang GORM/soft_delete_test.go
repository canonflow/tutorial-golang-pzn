package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type Todo struct {
	ID          int64          `gorm:"primaryKey;column:id;autoIncrement"`
	UserID      string         `gorm:"column:user_id"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (t *Todo) TableName() string {
	return "todos"
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserID:      "1",
		Title:       "Todo 1",
		Description: "Isi Todo 1",
	}

	// create
	result := db.Create(&todo)
	assert.Nil(t, result.Error)

	// delete
	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = db.Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))

	/*
		=== RUN   TestSoftDelete
		--- PASS: TestSoftDelete (0.00s)
		PASS
	*/
}

func TestUnscoped(t *testing.T) {
	var todo Todo

	result := db.Unscoped().First(&todo, "id = ?", "1")
	assert.Nil(t, result.Error)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))

	/*
		=== RUN TestUnscoped
		--- PASS: TestUnscoped (0.00s)
		PASS
	*/
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GuestBook struct {
	ID        int64  `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Message   string `gorm:"column:message"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt string `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (g *GuestBook) TableName() string {
	return "guest_books"
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)

	/*
		=== RUN   TestMigrator
		--- PASS: TestMigrator (0.02s)
		PASS
	*/
}

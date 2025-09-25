package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		u.ID = "user-" + time.Now().Format("20060102150405")
	}

	return nil
}

func TestUserHook(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName: "User 100",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.NotEqual(t, "", user.ID)

	/*
		=== RUN   TestUserHook
		--- PASS: TestUserHook (0.00s)
		PASS
	*/
}

package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

type Wallet struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}

type Address struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := db.Create(&wallet).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestCreateWallet
		--- PASS: TestCreateWallet (0.00s)
		PASS
	*/
}

func TestRetrieveRelation(t *testing.T) {
	var user User
	// err := db.Model(&User{}).Preload("Wallet").First(&user).Error
	err := db.Preload("Wallet").First(&user).Error

	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)

	/*
		=== RUN   TestRetrieveRelation
		--- PASS: TestRetrieveRelation (0.00s)
		PASS
	*/
}

func TestRetrieveRelationJoin(t *testing.T) {
	var users []User
	// Defaultnya pakai left join
	err := db.Model(&User{}).Joins("Wallet").Find(&users).Error

	assert.Nil(t, err)

	assert.Equal(t, 14, len(users))

	/*
		=== RUN   TestRetrieveRelationJoin
		--- PASS: TestRetrieveRelationJoin (0.00s)
		PASS
	*/
}

func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "20",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 20",
		},
		Wallet: Wallet{
			ID:      "20",
			UserId:  "20",
			Balance: 1000000,
		},
	}

	// Akan otomatis membuat data untuk tabel wallets juga
	err := db.Create(&user).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestAutoCreateUpdate
		--- PASS: TestAutoCreateUpdate (0.00s)
		PASS
	*/
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "21",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 21",
		},
		Wallet: Wallet{
			ID:      "21",
			UserId:  "21",
			Balance: 1000000,
		},
	}

	// Wallet tidak akan dibuat
	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestSkipAutoCreateUpdate
		--- PASS: TestSkipAutoCreateUpdate (0.00s)
		PASS
	*/
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "50",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 50",
		},
		Wallet: Wallet{
			ID:      "50",
			UserId:  "50",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				UserId:  "50",
				Address: "jalan A",
			},
			{
				UserId:  "50",
				Address: "Jalan B",
			},
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestUserAndAddresses
		--- PASS: TestUserAndAddresses (0.00s)
		PASS
	*/
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var userPreload []User

	err := db.Model(&User{}).
		Preload("Addresses"). // One to Many
		Joins("Wallet").      // One to One
		Find(&userPreload).Error

	assert.Nil(t, err)

	/*
		=== RUN   TestPreloadJoinOneToMany
		--- PASS: TestPreloadJoinOneToMany (0.00s)
		PASS
	*/
}

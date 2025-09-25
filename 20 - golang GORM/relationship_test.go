package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Wallet struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      *User     `gorm:"foreignKey:user_id;references:id"`

	// Gunakan pointer untuk menghindari cyclic dependency
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
	User      User      `gorm:"foreignKey:user_id;references:id"`
}

type Product struct {
	ID           string    `gorm:"primaryKey;column:id"`
	Name         string    `gorm:"column:name"`
	Price        int64     `gorm:"column:price"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUsers []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
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

func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Preload("User").Find(&addresses).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Joins("User").Find(&addresses).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestBelongsTo
		--- PASS: TestBelongsTo (0.00s)
		PASS
	*/
}

func TestBelongsToOneToOne(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet

	err := db.Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Joins("User").Find(&wallets).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestBelongsToOneToOne
		--- PASS: TestBelongsToOneToOne (0.00s)
		PASS
	*/
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		ID:    "P001",
		Name:  "Contoh Product",
		Price: 10000,
	}

	err := db.Create(&product).Error
	assert.Nil(t, err)

	// Create data into user_like_product table
	err = db.Table("user_like_product").
		Create(map[string]interface{}{
			"user_id":    "1",
			"product_id": "P001",
		}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").
		Create(map[string]interface{}{
			"user_id":    "2",
			"product_id": "P001",
		}).Error
	assert.Nil(t, err)

	/*
		=== RUN   TestCreateManyToMany
		--- PASS: TestCreateManyToMany (0.01s)
		PASS
	*/
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").First(&product, "id = ?", "P001").Error

	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))

	/*
		=== RUN   TestPreloadManyToMany
		--- PASS: TestPreloadManyToMany (0.00s)
		PASS
	*/
}

func TestAssociationFind(t *testing.T) {
	// Cari product tanpa users
	var product Product
	err := db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)
	// fmt.Println(product)

	// Cari user yang menyukai product tersebut
	var users []User
	err = db.Model(&product).
		Where("first_name LIKE ?", "User%").
		Association("LikedByUsers").
		Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	/*
		=== RUN   TestAssociationFind
		--- PASS: TestAssociationFind (0.00s)
		PASS
	*/
}

func TestAssociationAdd(t *testing.T) {
	// Mengambil User dengan ID 3
	var user User
	err := db.First(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	// Mengambil Product dengan ID P001
	var product Product
	err = db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	// Menambahkan Data ke tabel user_like_product dengan Product ID P001 dan User ID 3
	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)

	/*
		=== RUN   TestAssociationAdd
		--- PASS: TestAssociationAdd (0.01s)
		PASS
	*/
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.First(&user, "id = ?", "1").Error
		assert.Nil(t, err)
		fmt.Println(user)

		wallet := Wallet{
			ID:      "01",
			UserId:  "1",
			Balance: 2000000,
		}

		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})

	assert.Nil(t, err)

	/*
		=== RUN   TestAssociationReplace
		--- FAIL: TestAssociationReplace (0.01s)
		FAIL
	*/
}

func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.First(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	// Hapus user_id = 3 dan product_id = P001 di tabel user_like_product
	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)

	/*
		=== RUN   TestAssociationDelete
		--- PASS: TestAssociationDelete (0.01s)
		PASS
	*/
}

func TestAssocationClear(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	// Hapus semua data pada tabel user_like_product dengan product_id = P001
	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
	fmt.Println(product)

	/*
		=== RUN   TestAssocationClear
		--- PASS: TestAssocationClear (0.00s)
		PASS
	*/
}

func TestPreloadingWithCondition(t *testing.T) {
	var user User

	err := db.Preload("Wallet", "balance > ?", 1000000).First(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	fmt.Println(user)

	/*
		=== RUN   TestPreloadingWithCondition
		--- PASS: TestPreloadingWithCondition (0.00s)
		PASS
	*/
}

func TestNestedPreloading(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").Find(&wallet, "id = ? ", "1").Error
	assert.Nil(t, err)
	fmt.Println(wallet)

	/*
		=== RUN   TestNestedPreloading
		--- PASS: TestNestedPreloading (0.00s)
		PASS
	*/
}

func TestPreloadAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).First(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	fmt.Println(user)

	/*
		=== RUN   TestPreloadAll
		--- PASS: TestPreloadAll (0.00s)
		PASS
	*/
}

func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error // Using Left Join
	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))

	/*
		=== RUN   TestJoinQuery
		--- PASS: TestJoinQuery (0.01s)
		PASS
	*/
}

func TestJoinQueryCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users).Error

	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error // Alias menggunakan nama field
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	/*
		=== RUN   TestJoinQueryCondition
		--- PASS: TestJoinQueryCondition (0.00s)
		PASS
	*/
}

func TestCount(t *testing.T) {
	var count int64

	err := db.Model(&User{}).
		Joins("Wallet").
		Where("Wallet.balance > ?", 500000).
		Count(&count).Error

	assert.Nil(t, err)
	assert.Equal(t, int64(3), count)

	/*
		=== RUN   TestCount
		--- PASS: TestCount (0.00s)
		PASS
	*/
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestOtherAggregation(t *testing.T) {
	var result AggregationResult

	err := db.Model(&Wallet{}).
		Select("SUM(balance) as total_balance", "MIN(balance) as min_balance", "MAX(balance) as max_balance", "AVG(balance) as avg_balance").
		Take(&result).Error

	assert.Nil(t, err)
	assert.Equal(t, int64(4000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(2000000), result.MaxBalance)
	assert.Equal(t, float64(1333333.3333), result.AvgBalance)

	/*
		=== RUN   TestOtherAggregation
		--- PASS: TestOtherAggregation (0.00s)
		PASS
	*/
}

func TestGroupByHaving(t *testing.T) {
	var result []AggregationResult

	err := db.Model(&Wallet{}).
		Select("SUM(balance) as total_balance", "MIN(balance) as min_balance", "MAX(balance) as max_balance", "AVG(balance) as avg_balance").
		Joins("User").
		Group("User.id").
		Having("SUM(balance) > ?", 1000000).
		Find(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	/*
		=== RUN   TestGroupByHaving
		--- PASS: TestGroupByHaving (0.01s)
		PASS
	*/
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	err := db.WithContext(ctx).
		Find(&users).Error

	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))

	/*
		=== RUN   TestContext
		--- PASS: TestContext (0.00s)
		PASS
	*/
}

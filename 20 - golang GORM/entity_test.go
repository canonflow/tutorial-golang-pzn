package main

import "time"

type User struct {
	ID          string    `gorm:"primaryKey;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"`
}

type Name struct {
	FirstName  string `gorm:"first_name"`
	MiddleName string `gorm:"middle_name"`
	LastName   string `gorm:"last_name"`
}

func (u *User) TableName() string {
	return "users"
}

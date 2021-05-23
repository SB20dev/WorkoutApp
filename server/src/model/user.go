package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`
	Created   time.Time `json:"created"`
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func FetchUserByID(db *gorm.DB, id string) *User {
	var user User
	db.Where(&User{ID: id}).First(&user)
	return &user
}

func UpdateLastLogin(db *gorm.DB, id string) *gorm.DB {
	lastLogin := time.Now()
	return db.Model(&User{ID: id}).Update("last_login", lastLogin)
}

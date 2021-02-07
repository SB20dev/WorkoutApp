package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        string    `json:id`
	Password  string    `json:password`
	LastLogin time.Time `json:last_login`
	Created   time.Time `json:created`
}

func CreateUser(db *gorm.DB, user *User) *gorm.DB {
	return db.Create(user)
}

func FetchUserByID(db *gorm.DB, id string) *User {
	var user User
	db.Where(&User{ID: id}).First(&user, id)
	return &user
}

func UpdateLastLogin(db *gorm.DB, id string) *gorm.DB {
	lastLogin := time.Now()
	return db.Model(&User{ID: id}).Update("last_login", lastLogin)
}

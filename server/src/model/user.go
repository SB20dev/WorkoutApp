package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	SystemID  int64     `json:"system_id"`
	ID        string    `json:"id"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`
	Created   time.Time `json:"created"`
}

func CreateUser(db *gorm.DB, user *User) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			return err
		}
		err = DuplicateCommonRecords(db, user.SystemID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
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

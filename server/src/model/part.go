package model

import (
	"sync"

	"gorm.io/gorm"
)

type Part struct {
	ID     int    `json:"id"`
	Class  string `json:"class"`
	Detail string `json:"detail"`
}

var (
	once  sync.Once
	parts []Part
)

func GetParts(db *gorm.DB) ([]Part, error) {
	var err error
	once.Do(func() {
		parts := []Part{}
		err = db.Find(&parts).Error
	})
	if err != nil {
		return nil, err
	}
	return parts, nil
}

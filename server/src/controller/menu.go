package controller

import (
	"net/http"

	"gorm.io/gorm"
)

type MenuController struct {
	DB *gorm.DB
}

func (c *MenuController) Get(w http.ResponseWriter, r *http.Request, userID string) error {
	return nil
}

func (c *MenuController) Post(w http.ResponseWriter, r *http.Request) error {
	return nil
}

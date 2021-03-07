package controller

import (
	"net/http"
	"strconv"
	"workout/src/helper"
	"workout/src/model"

	"gorm.io/gorm"
)

type MenuController struct {
	DB *gorm.DB
}

func (c *MenuController) GetCount(w http.ResponseWriter, r *http.Request, userID string) error {
	count, err := model.FetchMenuCount(c.DB, userID)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to fetch count.")
	}

	rtn := map[string]int{
		"count": count,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *MenuController) GetByID(w http.ResponseWriter, r *http.Request, userID string) error {
	q := r.URL.Query()
	id, err := strconv.Atoi(q.Get("menu_id"))
	if err != nil {
		return helper.CreateHTTPError(http.StatusBadRequest, "query parameter is invalid. convertion failed.")
	}

	menu, err := model.FetchMenuByID(c.DB, userID, id)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to fetch menu.")
	}

	return helper.JSON(w, http.StatusOK, menu)
}

func (c *MenuController) GetPartially(w http.ResponseWriter, r *http.Request, userID string) error {
	// offset, num
	q := r.URL.Query()
	offset, offsetErr := strconv.Atoi(q.Get("offset"))
	num, numErr := strconv.Atoi(q.Get("num"))
	if offsetErr != nil || numErr != nil {
		return helper.CreateHTTPError(http.StatusBadRequest, "query parameter is invalid. convertion failed.")
	}

	menus, err := model.FetchMenus(c.DB, userID, offset, num)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, "failed to fetch menus.")
	}

	rtn := map[string](interface{}){
		"menus": menus,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

// TODO: 検索

func (c *MenuController) Post(w http.ResponseWriter, r *http.Request, userID string) error {
	return nil
}

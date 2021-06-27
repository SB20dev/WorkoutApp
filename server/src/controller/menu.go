package controller

import (
	"encoding/json"
	"errors"
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
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchMenuCountFailure)
	}

	rtn := map[string]int64{
		"count": count,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *MenuController) GetByID(w http.ResponseWriter, r *http.Request, userID string) error {
	q := r.URL.Query()
	id, err := strconv.Atoi(q.Get("menu_id"))
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidQueryParameter)
	}

	menu, err := model.FetchMenuByID(c.DB, id)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchMenuFailure)
	}

	return helper.JSON(w, http.StatusOK, menu)
}

func (c *MenuController) GetPartially(w http.ResponseWriter, r *http.Request, userID string) error {
	// offset, num
	q := r.URL.Query()
	offset, offsetErr := strconv.Atoi(q.Get("offset"))
	if offsetErr != nil {
		helper.LogError(r, offsetErr, nil)
	}
	num, numErr := strconv.Atoi(q.Get("num"))
	if numErr != nil {
		helper.LogError(r, numErr, nil)
	}
	if offsetErr != nil || numErr != nil {
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidQueryParameter)
	}

	menus, err := model.FetchMenus(c.DB, userID, offset, num)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchMenuFailure)
	}

	rtn := map[string](interface{}){
		"menus": menus,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

// TODO: 一時的な実装です
func (c *MenuController) Search(w http.ResponseWriter, r *http.Request, userID string) error {
	q := r.URL.Query()
	keyword := q.Get("keyword")
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidQueryParameter)
	}
	menus, err := model.SearchMenus(c.DB, userID, keyword, limit)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.FetchMenuFailure)
	}

	rtn := map[string](interface{}){
		"menus": menus,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (c *MenuController) Post(w http.ResponseWriter, r *http.Request, userID string) error {
	var body struct {
		Name  string `json:"name"`
		Parts []int  `json:"parts"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	} else if body.Name == "" || body.Parts == nil || len(body.Parts) == 0 {
		helper.LogError(r, errors.New("Menu name or parts is empty."), nil)
		return helper.CreateHTTPErrorWithCode(http.StatusBadRequest, helper.InvalidRequestBody)
	}

	menu := model.Menu{
		UserID: userID,
		Name:   body.Name,
	}

	err = model.CreateMenus(c.DB, &menu, body.Parts)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.CreationMenuFailure)
	}

	return nil
}

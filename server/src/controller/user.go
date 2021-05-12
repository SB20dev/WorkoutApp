package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"workout/src/helper"
	"workout/src/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (u *UserController) SignIn(w http.ResponseWriter, r *http.Request) error {
	var input model.User
	json.NewDecoder(r.Body).Decode(&input)

	// 入力のバリデーション
	if err := validateInputs(input); err != nil {
		return err
	}

	user := model.FetchUserByID(u.DB, input.ID)
	if user == nil {
		return helper.CreateHTTPError(http.StatusUnauthorized, "ID or password is not correct")
	}

	// パスワード照合
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return helper.CreateHTTPError(http.StatusUnauthorized, "ID or password is not correct")
	}

	// トークン生成
	tokenStr, err := helper.CreateToken(user.ID)
	if err != nil {
		return err
	}

	// 最終ログイン日時の更新
	result := model.UpdateLastLogin(u.DB, input.ID)
	if err := result.Error; err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, err.Error())
	}

	rtn := map[string]string{
		"token": tokenStr,
	}
	return helper.JSON(w, http.StatusOK, rtn)
}

func (u *UserController) SignUp(w http.ResponseWriter, r *http.Request) error {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	// 入力のバリデーション
	if err := validateInputs(user); err != nil {
		return err
	}

	user.Created = time.Now()
	// パスワード暗号化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, "password hash error")
	}
	user.Password = string(hash)

	// DB格納
	err = model.CreateUser(u.DB, &user)
	if err != nil {
		return helper.CreateHTTPError(http.StatusInternalServerError, err.Error())
	}

	return helper.JSON(w, http.StatusOK, nil)
}

func validateInputs(user model.User) error {

	addErrorStr := func(str string, addition string) string {
		if str != "" {
			return str + " " + addition
		}
		return addition
	}

	const (
		minLength int = 8
		maxLength int = 72
	)

	isInvalid := false
	errorStr := ""
	if user.ID == "" {
		errorStr += "ID is empty."
		isInvalid = true
	} else if len(user.ID) > maxLength || len(user.ID) < minLength {
		errorStr += fmt.Sprintf("length of ID must be from %d to %d.", minLength, maxLength)
		isInvalid = true
	}
	if user.Password == "" {
		errorStr = addErrorStr(errorStr, "Password is empty.")
		isInvalid = true
	} else if len(user.Password) > maxLength || len(user.Password) < minLength {
		errorStr = addErrorStr(errorStr, fmt.Sprintf("length of Password must be from %d to %d.", minLength, maxLength))
		isInvalid = true
	}

	if !isInvalid {
		return nil
	}
	return helper.CreateHTTPError(http.StatusBadRequest, errorStr)
}

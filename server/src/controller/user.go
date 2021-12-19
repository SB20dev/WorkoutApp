package controller

import (
	"encoding/json"
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
		return helper.CreateHTTPErrorWithCode(http.StatusUnauthorized, helper.IncorrectUserIdOrPassword)
	}

	// パスワード照合
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return helper.CreateHTTPErrorWithCode(http.StatusUnauthorized, helper.IncorrectUserIdOrPassword)
	}

	// トークン生成
	tokenStr, err := helper.CreateToken(user.SystemID)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UserError)
	}

	// 最終ログイン日時の更新
	result := model.UpdateLastLogin(u.DB, input.ID)
	if err := result.Error; err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UserError)
	}

	// tokenをcookieに設定
	cookie := &http.Cookie{
		Name:  "token",
		Value: tokenStr,
	}
	http.SetCookie(w, cookie)

	return helper.JSON(w, http.StatusOK, nil)
}

func (u *UserController) SignUp(w http.ResponseWriter, r *http.Request) error {
	var input model.User
	json.NewDecoder(r.Body).Decode(&input)

	helper.Logf(r, "%v", input)

	// 入力のバリデーション
	if err := validateInputs(input); err != nil {
		return err
	}

	// IDの重複チェック
	user := model.FetchUserByID(u.DB, input.ID)
	if user == nil {
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.DuplicatedUserId)
	}

	input.Created = time.Now()
	// パスワード暗号化
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UserError)
	}
	input.Password = string(hash)

	// DB格納
	err = model.CreateUser(u.DB, &input)
	if err != nil {
		helper.LogError(r, err, nil)
		return helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UserError)
	}

	return helper.JSON(w, http.StatusOK, nil)
}

func validateInputs(user model.User) error {

	const (
		minLength int = 8
		maxLength int = 72
	)

	errorCodes := []int{}
	if len(user.ID) > maxLength || len(user.ID) < minLength {
		errorCodes = append(errorCodes, helper.InvalidUserId)
	}

	if len(user.Password) > maxLength || len(user.Password) < minLength {
		errorCodes = append(errorCodes, helper.InvalidUserPassword)
	}

	if len(errorCodes) > 0 {
		return helper.CreateHTTPErrorWithCodes(http.StatusBadRequest, errorCodes)
	}
	return nil
}

package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"workout/src/model"
)

type userTestData struct {
	userID       string
	password     string
	expectedCode int
	expectedBody string
}

func TestSignUp(t *testing.T) {
	tx := database.Begin()

	// 正常ルート
	successData := userTestData{
		userID:       "hogehoge12",
		password:     "fugafuga12",
		expectedCode: 200,
		expectedBody: "{}",
	}
	body := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, successData.userID, successData.password))
	req, err := http.NewRequest("POST", "/api/user/signup", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// リクエスト実行
	router.ServeHTTP(rr, req)

	// 照合
	// ステータスコード確認
	if status := rr.Code; status != successData.expectedCode {
		t.Errorf("wrong status code: got %v want %v",
			status, successData.expectedCode)
	}
	// レスポンスボディ確認
	if respBody := rr.Body.String(); strings.TrimSpace(respBody) != successData.expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			respBody, successData.expectedBody)
	}
	// DB確認
	var user model.User
	err = database.Where(&model.User{ID: successData.userID}).First(&user).Error
	if err != nil {
		t.Error("insertion failed.")
	}

	// エラールート
	errorDatas := []userTestData{
		{
			userID:       "",
			password:     "fugafuga12",
			expectedCode: 400,
			expectedBody: "Internal Server error : ID is empty.",
		},
		{
			userID:       "hoge",
			password:     "fugafuga12",
			expectedCode: 400,
			expectedBody: "Internal Server error : length of ID must be from 8 to 72.",
		},
		{
			userID:       "hogehoge12",
			password:     "",
			expectedCode: 400,
			expectedBody: "Internal Server error : Password is empty.",
		},
		{
			userID:       "hogehoge12",
			password:     "fuga",
			expectedCode: 400,
			expectedBody: "Internal Server error : length of Password must be from 8 to 72.",
		},
	}
	for _, errorData := range errorDatas {
		body := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, errorData.userID, errorData.password))
		req, err := http.NewRequest("POST", "/api/user/signup", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		// リクエスト実行
		router.ServeHTTP(rr, req)

		// 照合
		// ステータスコード確認
		if status := rr.Code; status != errorData.expectedCode {
			t.Errorf("wrong status code: got %v want %v",
				status, errorData.expectedCode)
		}
		// レスポンスボディ確認
		if respBody := rr.Body.String(); strings.TrimSpace(respBody) != strings.TrimSpace(errorData.expectedBody) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				respBody, errorData.expectedBody)
		}
	}

	tx.Rollback()
}

// func TestSignIn(t *testing.T) {
// 	tx := database.Begin()

// 	body := bytes.NewBufferString("{id:hogehoge12, password:fugafuga12}")
// 	req, err := http.NewRequest("POST", "/api/user/signup", body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()

// 	// リクエスト実行
// 	router.ServeHTTP(rr, req)

// 	var user model.User
// 	database.Where(&model.User{ID: "hogehoge12"}).First(&user)

// 	tx.Rollback()
// }

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"workout/src/handler"
	"workout/src/model"
)

type userTestData struct {
	userID       string
	password     string
	expectedCode int
	expectedBody string
}

func TestSignUpSuccess(t *testing.T) {
	tx := database.Begin()
	router := handler.GetRouter(tx)

	// 正常ルート
	var successData = userTestData{
		userID:       "hogehoge12",
		password:     "fugafuga12",
		expectedCode: http.StatusOK,
		expectedBody: "{}",
	}
	body := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, successData.userID, successData.password))
	req, err := http.NewRequest("POST", "/api/user/signup", body)
	req.Header.Set("Content-Type", "application/json")
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
	if respBody := rr.Body.String(); !checkJsonEquality(successData.expectedBody, respBody) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			respBody, successData.expectedBody)
	}

	// DB確認
	var user model.User
	err = tx.Where(&model.User{ID: successData.userID}).First(&user).Error
	if err != nil {
		t.Error("insertion failed.")
	}

	tx.Rollback()
}

func TestSignUpFailure(t *testing.T) {
	tx := database.Begin()
	router := handler.GetRouter(tx)

	// エラールート
	var errorDatas = []userTestData{
		{
			userID:       "",
			password:     "fugafuga12",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "ID is empty." }`,
		},
		{
			userID:       "hoge",
			password:     "fugafuga12",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "length of ID must be from 8 to 72." }`,
		},
		{
			userID:       "hogehoge12",
			password:     "",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "Password is empty." }`,
		},
		{
			userID:       "hogehoge12",
			password:     "fuga",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "length of Password must be from 8 to 72." }`,
		},
	}

	for _, errorData := range errorDatas {
		body := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, errorData.userID, errorData.password))
		req, err := http.NewRequest("POST", "/api/user/signup", body)
		req.Header.Set("Content-Type", "application/json")
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
		if respBody := rr.Body.String(); !checkJsonEquality(errorData.expectedBody, respBody) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				respBody, errorData.expectedBody)
		}
	}

	tx.Rollback()
}

func TestSignInSuccess(t *testing.T) {
	tx := database.Begin()
	router := handler.GetRouter(tx)

	// 準備
	// ユーザ登録
	var successData = userTestData{
		userID:       "hogehoge12",
		password:     "fugafuga12",
		expectedCode: http.StatusOK,
	}
	reqBody := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, successData.userID, successData.password))
	req, err := http.NewRequest("POST", "/api/user/signup", reqBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// リクエスト実行
	router.ServeHTTP(rr, req)

	// リクエスト実行
	reqBody = bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, successData.userID, successData.password))
	req, err = http.NewRequest("POST", "/api/user/signin", reqBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// レスポンス確認
	// ステータスコード確認
	if status := rr.Code; status != successData.expectedCode {
		t.Errorf("wrong status code: got %v want %v",
			status, successData.expectedCode)
	}
	// 認証確認
	var respBody map[string]string
	json.Unmarshal([]byte(rr.Body.String()), &respBody)

	req, err = http.NewRequest("GET", "/api/user/checkauth", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", respBody["token"]))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if strings.TrimSpace(rr.Body.String()) != "authorized" {
		t.Errorf("Authorization failed.")
	}

	tx.Rollback()
}

func TestSignInFailure(t *testing.T) {
	tx := database.Begin()
	router := handler.GetRouter(tx)

	// 準備
	// ユーザ登録
	var successData = userTestData{
		userID:       "hogehoge12",
		password:     "fugafuga12",
		expectedCode: http.StatusOK,
	}
	reqBody := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, successData.userID, successData.password))
	req, err := http.NewRequest("POST", "/api/user/signup", reqBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// テストデータ
	var errorDatas = []userTestData{
		{
			userID:       "",
			password:     "fugafuga12",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "ID is empty." }`,
		},
		{
			userID:       "hoge",
			password:     "fugafuga12",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "length of ID must be from 8 to 72." }`,
		},
		{
			userID:       "hogehoge12",
			password:     "",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "Password is empty." }`,
		},
		{
			userID:       "hogehoge12",
			password:     "fuga",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{ "error" : "length of Password must be from 8 to 72." }`,
		},
		{
			userID:       "hogehoge123",
			password:     "fugafuga12",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{ "error" : "ID or password is not correct." }`,
		},
		{
			userID:       "hogehoge12",
			password:     "fugafuga123",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{ "error" : "ID or password is not correct." }`,
		},
	}
	for _, errorData := range errorDatas {
		// リクエスト実行
		body := bytes.NewBufferString(fmt.Sprintf(`{"id":"%s", "password":"%s"}`, errorData.userID, errorData.password))
		req, err := http.NewRequest("POST", "/api/user/signin", body)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// 照合
		// ステータスコード確認
		if status := rr.Code; status != errorData.expectedCode {
			t.Errorf("wrong status code: got %v want %v",
				status, errorData.expectedCode)
		}
		// レスポンスボディ確認
		if respBody := rr.Body.String(); !checkJsonEquality(errorData.expectedBody, respBody) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				respBody, errorData.expectedBody)
		}
	}

	tx.Rollback()
}

func checkJsonEquality(jsonStr1, jsonStr2 string) bool {
	// Unmarshal
	var json1, json2 map[string]string
	json.Unmarshal([]byte(jsonStr1), &json1)
	json.Unmarshal([]byte(jsonStr2), &json2)

	// 等価判定
	equality := true
	for key, value := range json1 {
		if value != json2[key] {
			equality = false
			break
		}
	}
	return equality
}

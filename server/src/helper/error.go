package helper

import (
	"errors"
	"fmt"
	"strings"
)

type HTTPError struct {
	StatusCode int
	ErrorCodes []int
	Err        error
}

func (err *HTTPError) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("status: %d, reason: %s", err.StatusCode, err.Err.Error())
	}
	return fmt.Sprintf("Status: %d", err.StatusCode)
}

func CreateHTTPErrorWithCode(status, errorCode int) *HTTPError {
	return &HTTPError{StatusCode: status, ErrorCodes: []int{errorCode}, Err: errors.New(errorMap[errorCode])}
}

func CreateHTTPErrorWithCodes(status int, errorCodes []int) *HTTPError {
	return &HTTPError{StatusCode: status, ErrorCodes: errorCodes, Err: errors.New(getErrorStr(errorCodes))}
}

func CreateHTTPErrorWithMessage(status int, errorStr string) *HTTPError {
	return &HTTPError{StatusCode: status, ErrorCodes: []int{UnrecognizedError}, Err: errors.New(errorStr)}
}

var errorMap = map[int]string{
	// 汎用
	InvalidQueryParameter: "Query parameter is invalid.",
	InvalidRequestBody:    "Request body is invalid.",
	// ユーザー系
	UserError:                 "Error about users.",
	InvalidUserId:             "User ID is invalid.",
	InvalidUserPassword:       "User Password is invalid.",
	DuplicatedUserId:          "User ID is duplicated.",
	IncorrectUserIdOrPassword: "User ID or password is not correct.",
	// メニュー系
	MenuError:             "Error about menus.",
	FetchMenuCountFailure: "Failed to fetch Menu count.",
	FetchMenuFailure:      "Failed to fetch Menus.",
	CreationMenuFailure:   "Failed to create Menus.",
	// コミットメント系
	CommitmentError:                  "Error about commitments.",
	FetchTotalCommitmentScoreFailure: "Failed to fetch total commitment score.",
	FetchCommitmentCountFailure:      "Failed to fetch commitment count.",
	FetchCommitmentHistoryFailure:    "Failed to fetch commitment histories.",
	FetchCommitmentDetailFailure:     "Failed to fetch commitment details.",
}

const (
	// 汎用
	UnrecognizedError     = 100
	InvalidQueryParameter = 101
	InvalidRequestBody    = 102
	// ユーザー系
	UserError                 = 200
	InvalidUserId             = 201
	InvalidUserPassword       = 202
	DuplicatedUserId          = 203
	IncorrectUserIdOrPassword = 204
	// メニュー系
	MenuError             = 300
	FetchMenuCountFailure = 301
	FetchMenuFailure      = 302
	CreationMenuFailure   = 303
	// コミットメント系
	CommitmentError                  = 400
	FetchTotalCommitmentScoreFailure = 401
	FetchCommitmentCountFailure      = 402
	FetchCommitmentHistoryFailure    = 403
	FetchCommitmentDetailFailure     = 404
	CreationCommitmentFailure        = 405
)

func getErrorStr(errorCodes []int) string {
	errorStr := make([]string, len(errorCodes))
	for idx, code := range errorCodes {
		value, exists := errorMap[code]
		if exists {
			errorStr[idx] = value
		}
	}
	return strings.Join(errorStr, "\n")
}

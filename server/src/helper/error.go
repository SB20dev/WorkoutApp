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
	return &HTTPError{StatusCode: status, ErrorCodes: []int{errorCode}, Err: errors.New(ErrorMap[errorCode])}
}

func CreateHTTPErrorWithCodes(status int, errorCodes []int) *HTTPError {
	return &HTTPError{StatusCode: status, ErrorCodes: errorCodes, Err: errors.New(GetErrorStr(errorCodes))}
}

func CreateHTTPErrorWithMessage(status int, errorStr string) *HTTPError {
	return &HTTPError{StatusCode: status, ErrorCodes: []int{UnrecognizedError}, Err: errors.New(errorStr)}
}

var ErrorMap = map[int]string{
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
	MenuError:              "Error about menus.",
	FetchMenusCountFailure: "Failed to fetch menus count.",
	FetchMenuFailure:       "Failed to fetch menus.",
	CreateMenuFailure:      "Failed to create menus.",
	// コミットメント系
	CommitmentError:                  "Error about commitments.",
	FetchTotalCommitmentScoreFailure: "Failed to fetch total commitment score.",
	FetchCommitmentsCountFailure:     "Failed to fetch commitments count.",
	FetchCommitmentsFailure:          "Failed to fetch commitment histories.",
	FetchCommitmentDetailFailure:     "Failed to fetch commitment details.",
	CreateCommitmentFailure:          "Failed to create commitment.",
	// パーツ系
	PartError:              "Error about parts.",
	FetchPartsCountFailure: "Failed to fetch parts count.",
	FetchPartsFailure:      "Failed to fetch parts.",
	CreatePartFailure:      "Failed to create part.",
	UpdatePartFailure:      "Failed to update part.",
	DeletePartFailure:      "Failed to delete part.",
	FetchClassesFailure:    "Failed to fetch classes.",
	CreateClassFailure:     "Failed to create class.",
	UpdateClassFailure:     "Failed to update class.",
	DeleteClassFailure:     "Failed to delete class.",
	FetchStatusFailure:     "Failed to fetch status.",
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
	MenuError              = 300
	FetchMenusCountFailure = 301
	FetchMenuFailure       = 302
	CreateMenuFailure      = 303
	// コミットメント系
	CommitmentError                  = 400
	FetchTotalCommitmentScoreFailure = 401
	FetchCommitmentsCountFailure     = 402
	FetchCommitmentsFailure          = 403
	FetchCommitmentDetailFailure     = 404
	CreateCommitmentFailure          = 405
	// パーツ系
	PartError              = 500
	FetchPartsCountFailure = 501
	FetchPartsFailure      = 502
	CreatePartFailure      = 503
	UpdatePartFailure      = 504
	DeletePartFailure      = 505
	FetchClassesFailure    = 506
	CreateClassFailure     = 507
	UpdateClassFailure     = 508
	DeleteClassFailure     = 509
	FetchStatusFailure     = 510
)

func GetErrorStr(errorCodes []int) string {
	errorStr := make([]string, len(errorCodes))
	for idx, code := range errorCodes {
		value, exists := ErrorMap[code]
		if exists {
			errorStr[idx] = value
		}
	}
	return strings.Join(errorStr, "\n")
}

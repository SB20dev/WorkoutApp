package test

import (
	"encoding/json"
	"fmt"
	"workout/src/helper"
)

const ErrorResponseBodyFormat = `{"error":"status: %d, reason: %s","errorCodes":%v}`

func GetErrorResponseBody(statusCode int, errorCodes []int) string {
	return fmt.Sprintf(ErrorResponseBodyFormat, statusCode, helper.GetErrorStr(errorCodes), errorCodes)
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

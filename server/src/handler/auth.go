package handler

import (
	"net/http"
	"os"
	"strconv"
	"workout/src/helper"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNATURE")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func AuthHandler(fn func(w http.ResponseWriter, r *http.Request, userID int64) error) http.Handler {
	funcWithUserID := func(w http.ResponseWriter, r *http.Request) error {
		strSystemUserID, ok := helper.GetClaim(r, "id").(string)
		if !ok {
			return helper.CreateHTTPErrorWithMessage(http.StatusInternalServerError, "failed to read user id from context")
		}
		systemUserID, err := strconv.ParseInt(strSystemUserID, 10, 64)
		if err != nil {
			return helper.CreateHTTPErrorWithMessage(http.StatusInternalServerError, "failed to parse user id from context")
		}
		return fn(w, r, systemUserID)
	}
	return jwtMiddleware.Handler(Handler(funcWithUserID))
}

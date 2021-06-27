package handler

import (
	"net/http"
	"os"
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

func AuthHandler(fn func(w http.ResponseWriter, r *http.Request, userID string) error) http.Handler {
	funcWithUserID := func(w http.ResponseWriter, r *http.Request) error {
		userID, ok := helper.GetClaim(r, "id").(string)
		if !ok {
			return helper.CreateHTTPErrorWithMessage(http.StatusInternalServerError, "failed to read user id from context")
		}
		return fn(w, r, userID)
	}
	return jwtMiddleware.Handler(Handler(funcWithUserID))
}

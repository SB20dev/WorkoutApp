package helper

import (
	"WorkoutApp/server/src/model"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

func CreateToken(user *model.User) (string, error) {

	// Token を作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// registered claims
		"iss": "sb20",
		"sub": "workout_app",
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		// private claims
		"id": user.ID,
	})

	signature := os.Getenv("SIGNATURE")
	tokenString, err := token.SignedString([]byte(signature))

	if err != nil {
		return "", CreateHTTPError(http.StatusInternalServerError, err.Error())
	}

	return tokenString, nil
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNATURE")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func AuthHandler(fn handlerFunc) http.Handler {
	return jwtMiddleware.Handler(Handler(fn))
}

func GetClaim(r *http.Request, key string) interface{} {
	claims := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims[key]
}
package helper

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

func JSON(w http.ResponseWriter, code int, data interface{}) error {
	w.WriteHeader(code)
	if data == nil {
		data = map[string]interface{}{}
	}
	return json.NewEncoder(w).Encode(data)
}

func CreateToken(id string) (string, error) {

	// Token を作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// registered claims
		"iss": "sb20",
		"sub": "workout_app",
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		// private claims
		"id": id,
	})

	signature := os.Getenv("SIGNATURE")
	tokenString, err := token.SignedString([]byte(signature))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetClaim(r *http.Request, key string) interface{} {
	claims := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims[key]
}

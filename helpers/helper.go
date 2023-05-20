package helpers

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	SALT       = 8
	SECRET_KEY = "thisissecret"
)

func HashPassword(pass string) string {
	password := []byte(pass)

	hash, _ := bcrypt.GenerateFromPassword(password, SALT)

	return string(hash)
}

func ComparePassword(hashed, password []byte) bool {
	h, p := []byte(hashed), []byte(password)

	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}

func GenerateToken(id int, username string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))

	return signedToken
}

func VerifyToken(tokenString string) (interface{}, error) {
	errResponse := errors.New("Token-Invalid")

	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(SECRET_KEY), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

func FetchAPI(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

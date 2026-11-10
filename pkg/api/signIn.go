package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type UserPassword struct {
	Password string `json:"password"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var p UserPassword
	todo_password := os.Getenv("TODO_PASSWORD")

	fmt.Println(todo_password, "sect")
	if todo_password == "" {

		writeJson(w, map[string]string{
			"error": "Пароль не установлен",
		}, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {

		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if p.Password == "" {

		writeJson(w, map[string]string{
			"error": "Неверный пароль",
		}, http.StatusBadRequest)
		return
	}
	if p.Password != todo_password {

		writeJson(w, map[string]string{
			"error": "Неверный пароль",
		}, http.StatusBadRequest)
		return
	}
	token, err := GenerateToken(p.Password)
	if err != nil {

		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]string{
		"token": token,
	}, http.StatusOK)
}

func GenerateToken(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	passwordChecksum := hex.EncodeToString(hash[:])

	claims := jwt.MapClaims{
		"pwd": passwordChecksum,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	fmt.Println(os.Getenv("TODO_JWT_SECRET"), "123213123")
	return token.SignedString([]byte(os.Getenv("TODO_JWT_SECRET")))
}

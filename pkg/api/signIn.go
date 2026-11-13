package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type UserPassword struct {
	Password string `json:"password"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJson(w, "method not allowd", http.StatusMethodNotAllowed)
	}

	var p UserPassword
	todo_password := Cfg.TODO_PASSWORD

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

	return token.SignedString([]byte(Cfg.TODO_JWT_SECRET))
}

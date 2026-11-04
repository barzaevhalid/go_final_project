package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "Пароль не установлен",
		})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if p.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "Неверный пароль",
		})
		return
	}
	if p.Password != todo_password {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "Неверный пароль",
		})
		return
	}
	token, err := GenerateToken(p.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJson(w, map[string]string{
		"token": token,
	})
}

func GenerateToken(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	passwordChecksum := hex.EncodeToString(hash[:])

	claims := jwt.MapClaims{
		"pwd": passwordChecksum,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

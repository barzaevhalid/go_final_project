package api

import (
	"net/http"
)

func auth(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pass := Cfg.TODO_PASSWORD
		if len(pass) > 0 {
			var jwt string

			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}
			var valid bool

			valid, err = ValidateToken(jwt)

			if err != nil || !valid {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

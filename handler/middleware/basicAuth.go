package middleware

import (
	"fmt"
	"net/http"
	"os"
)

func BasicAuth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		basicAuthUserID := os.Getenv("BASIC_AUTH_USER_ID")
		basicAuthPassword := os.Getenv("BASIC_AUTH_PASSWORD")
		r.SetBasicAuth("", "")
		userID, password, ok := r.BasicAuth()
		fmt.Println(userID)
		fmt.Println(basicAuthUserID)
		fmt.Println(password)
		fmt.Println(basicAuthPassword)
		if !ok || userID == "" || password == "" || userID != basicAuthUserID || password != basicAuthPassword {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

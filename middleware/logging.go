package middleware

import (
	"fmt"
	"net/http"
)

func LoggingUri(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

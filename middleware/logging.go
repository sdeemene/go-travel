package middleware

import (
	"fmt"
	"net/http"
)

func LoggingUri(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr, r.RequestURI, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

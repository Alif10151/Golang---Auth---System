package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)

		duration := time.Since(start)
		fmt.Printf("%s %s took %v\n", r.Method, r.URL.Path, duration)
	}
}

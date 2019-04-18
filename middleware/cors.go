package middleware

import "net/http"

// CORS func Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

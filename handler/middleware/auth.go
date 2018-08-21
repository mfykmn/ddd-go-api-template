package middleware

import (
	"log"
	"net/http"
)

// AuthMiddleware 前処理として認証を行う
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("exec AuthMiddleware")
		// Auth
		next.ServeHTTP(w, r)
	})
}

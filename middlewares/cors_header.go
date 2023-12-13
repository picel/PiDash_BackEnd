package middlewares

import (
	"net/http"
)

// CorsHeader adds cors headers for the mux router
func CorsHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}

package rest

import "net/http"

func (s *server) Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Add("Access-Control-Allow-Origin", "*")

		(w).Header().Set("Content-Type", "text/html; charset=utf-8")

		h.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

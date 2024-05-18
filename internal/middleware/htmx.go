package middleware

import (
	"net/http"
)

func CheckHtmx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.Header.Get("Hx-Request") == "true" {
			setRequestContext(ctx, r, "hx-request", true)
			w.Header().Set("Cache-Control", "no-store, max-age=0")
		}
		next.ServeHTTP(w, r)
	})
}

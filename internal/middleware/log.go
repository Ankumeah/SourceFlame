package middleware

import (
	"log"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf(
				"%v %v %v\n",
				r.RemoteAddr, r.Method, r.URL.Path,
			)

			next.ServeHTTP(w, r)
		},
	)
}

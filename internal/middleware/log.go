package middleware

import (
	"context"
	"log"
	"net/http"
)

const ErrorsKey = Key("errors")

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf(
				"%s %s %s\n",
				r.RemoteAddr, r.Method, r.URL.Path,
			)

			ctx := context.WithValue(r.Context(), ErrorsKey, []map[string]any{})

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

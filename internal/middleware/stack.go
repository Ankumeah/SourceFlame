package middleware

import (
	"net/http"
	"slices"
)

type middleware func(http.Handler) http.Handler

func Stack(xs ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		for _, x := range slices.Backward(xs) {
			next = x(next)
		}

		return next
	}
}

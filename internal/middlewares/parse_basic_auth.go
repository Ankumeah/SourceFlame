package middlewares

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/Ankumeah/SourceFlame/internal/utils"
)

const UserKey = Key("user")
const PassKey = Key("pass")

func ParseBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			auth := strings.TrimSpace(r.Header.Get("Authorization"))
			if len(auth) <= 0 {
				utils.RespMsg(
					"Authorization not provided",
					http.StatusInternalServerError,
					w,
				)
				return
			}

			authFields := strings.Fields(auth)
			if len(authFields) != 2 {
				utils.RespMsg(
					"Incorrect amount of segements Authorization",
					http.StatusInternalServerError,
					w,
				)
				return
			}

			tokenType := authFields[0]
			token := authFields[1]

			if tokenType != "Basic" {
				utils.RespMsg(
					"Unsupported auth type",
					http.StatusBadRequest,
					w,
				)
				return
			}

			decodedToken, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				utils.RespMsg(
					"Invalid base64 auth",
					http.StatusBadRequest,
					w,
				)
				return
			}

			username, pass, ok := strings.Cut(string(decodedToken), ":")
			if !ok {
				utils.RespMsg(
					"Invalid auth format",
					http.StatusBadRequest,
					w,
				)
				return
			}

			if !utils.IsValidName(username) {
				utils.RespMsg(
					"Username not allowed",
					http.StatusBadRequest,
					w,
				)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, username)
			ctx = context.WithValue(ctx, PassKey, pass)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

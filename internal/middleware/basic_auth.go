package middleware

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/Ankumeah/SourceFlame/internal/utils"
)

const UserKey = Key("user")
const PassKey = Key("pass")

func BasicAuth(next http.Handler) http.Handler {
	isUsernameValid := regexp.MustCompilePOSIX("^[a-zA-Z0-9_-]{4,}$")
	isUsernameValid.Longest()

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			errs, ok := r.Context().Value(ErrorsKey).([]map[string]any)
			if !ok {
				log.Printf("%s %s: error map not set", r.Method, r.URL.Path)
				w.WriteHeader(http.StatusInternalServerError)
			}

			auth := strings.TrimSpace(r.Header.Get("Authorization"))
			if len(auth) <= 0 {
				utils.RespMsg(
					"Authorization not provided",
					http.StatusInternalServerError,
					errs, w,
				)
				return
			}

			authFields := strings.Fields(auth)
			if len(authFields) != 2 {
				utils.RespMsg(
					"Incorrect amount of segements Authorization",
					http.StatusInternalServerError,
					errs, w,
				)
				return
			}

			tokenType := authFields[0]
			token := authFields[1]

			if tokenType != "Basic" {
				utils.RespMsg(
					"Unsupported auth type",
					http.StatusBadRequest,
					errs, w,
				)
				return
			}

			decodedToken, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				utils.RespMsg(
					"Invalid base64 auth",
					http.StatusBadRequest,
					errs, w,
				)
				return
			}

			username, pass, ok := strings.Cut(string(decodedToken), ":")
			if !ok {
				utils.RespMsg(
					"Invalid auth format",
					http.StatusBadRequest,
					errs, w,
				)
				return
			}

			if username != string(isUsernameValid.Find([]byte(username))) {
				utils.RespMsg(
					"Username not allowed",
					http.StatusBadRequest,
					errs, w,
				)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, username)
			ctx = context.WithValue(ctx, PassKey, pass)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

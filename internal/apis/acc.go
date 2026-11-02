package apis

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/middlewares"
	"github.com/Ankumeah/SourceFlame/internal/utils"
)

func acc(a *app.App, root Route) {
	parsed := root.Sub("/acc", middlewares.ParseBasicAuth)
	authed := root.Sub("/acc", func(h http.Handler) http.Handler { return middlewares.BasicAuth(h, a) })

	parsed.Base(
		"POST",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user, ok := ctx.Value(middlewares.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), w)
				return
			}

			pass, ok := ctx.Value(middlewares.PassKey).(string)
			if !ok {
				utils.RespErr(errors.New("pass set by middleware not string"), w)
				return
			}

			hash, err := a.Hasher.GenerateHash([]byte(pass), nil)
			if err != nil {
				utils.RespErr(err, w)
				return
			}

			if !handle(a.DB.AddUser(ctx, user, hash), w) {
				return
			}

			w.WriteHeader(http.StatusOK)
		},
	)

	authed.Base(
		"PATCH",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user, ok := ctx.Value(middlewares.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), w)
				return
			}

			request := struct {
				Email string `json:"email"`
			}{}

			if r.Body == nil {
				utils.RespMsg(
					"No json body provided",
					http.StatusBadRequest,
					w,
				)
				return
			}
			defer r.Body.Close()

			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				utils.RespMsg(
					"Invalid json body",
					http.StatusBadRequest,
					w,
				)
				return
			}

			if !handle(a.DB.ChangeEmail(ctx, user, request.Email), w) {
				return
			}

			w.WriteHeader(http.StatusOK)
		},
	)

	authed.Base(
		"DELETE",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user, ok := ctx.Value(middlewares.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), w)
				return
			}

			if !handle(a.DB.DeleteUser(ctx, user), w) {
				return
			}

			w.WriteHeader(http.StatusOK)
		},
	)
}

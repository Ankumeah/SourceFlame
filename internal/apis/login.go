package apis

import (
	"errors"
	"log"
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/db"
	"github.com/Ankumeah/SourceFlame/internal/middleware"
	"github.com/Ankumeah/SourceFlame/internal/utils"
)

func login(a *app.App, mux *http.ServeMux) http.Handler {
	route := subRoute("/login", mux)
	handler := middleware.Stack(
		middleware.BasicAuth,
	)(route)

	route.HandleFunc(
		"POST /",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			errs, ok := ctx.Value(middleware.ErrorsKey).([]map[string]any)
			if !ok {
				log.Printf("%s %s: error map not set", r.Method, r.URL.Path)
				w.WriteHeader(http.StatusInternalServerError)
			}

			user, ok := ctx.Value(middleware.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), errs, w)
				return
			}

			pass, ok := ctx.Value(middleware.PassKey).(string)
			if !ok {
				utils.RespErr(errors.New("pass set by middleware not string"), errs, w)
				return
			}

			hash, err := a.Hasher.GenerateHash([]byte(pass), nil)
			if err != nil {
				utils.RespErr(err, errs, w)
				return
			}

			if err := a.DB.AddUser(ctx, db.User{
				Name: user,
			}, hash); !errors.Is(err, db.ErrExistsUser) && err != nil {
				utils.RespErr(err, errs, w)
				return
			}

			// TODO: Add session
		},
	)

	return handler
}

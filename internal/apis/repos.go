package apis

import (
	"errors"
	"net/http"
	"path/filepath"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/db"
	e "github.com/Ankumeah/SourceFlame/internal/errs"
	"github.com/Ankumeah/SourceFlame/internal/middlewares"
	"github.com/Ankumeah/SourceFlame/internal/utils"
)

func repos(a *app.App, root Route) {
	route := root.Sub("/repos/{repo}", func(h http.Handler) http.Handler { return middlewares.BasicAuth(h, a) })

	route.Base(
		"POST",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user, ok := ctx.Value(middlewares.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), w)
				return
			}

			repo := r.PathValue("repo")
			if !utils.IsValidName(repo) {
				utils.RespMsg(
					"Invalid repo name",
					http.StatusBadRequest,
					w,
				)
				return
			}

			if err := a.DB.AddRepo(
				ctx,
				db.Repo{
					Owner: user,
					Name:  repo,
				},
			); errors.Is(err, e.ErrInvalidUser) {
				utils.RespMsg(
					err.Error(),
					http.StatusNotFound,
					w,
				)
				return
			} else if errors.Is(err, e.ErrExistsRepo) {
				utils.RespMsg(
					err.Error(),
					http.StatusConflict,
					w,
				)
				return
			} else if err != nil {
				utils.RespErr(err, w)
				return
			}

			if err := a.GitController.CreateRepo(
				ctx, filepath.Join(user, repo),
			); err != nil {
				utils.RespErr(err, w)
				a.GitController.DeleteRepo(filepath.Join(user, repo))
				return
			}

			w.WriteHeader(http.StatusOK)
		},
	)

	route.Base(
		"DELETE",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user, ok := ctx.Value(middlewares.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), w)
				return
			}

			repo := r.PathValue("repo")
			if !utils.IsValidName(repo) {
				utils.RespMsg(
					"Invalid repo name",
					http.StatusBadRequest,
					w,
				)
				return
			}

			if err := a.DB.DeleteRepo(
				ctx,
				db.Repo{
					Owner: user,
					Name:  repo,
				},
			); errors.Is(err, e.ErrInvalidRepo) {
				utils.RespMsg(
					err.Error(),
					http.StatusNotFound,
					w,
				)
				return
			} else if err != nil {
				utils.RespErr(err, w)
				return
			}

			a.GitController.DeleteRepo(filepath.Join(user, repo))

			w.WriteHeader(http.StatusOK)
		},
	)
}

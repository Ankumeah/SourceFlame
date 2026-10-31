package apis

import (
	"errors"
	"net/http"
	"path/filepath"

	"github.com/Ankumeah/SourceFlame/internal/app"
	e "github.com/Ankumeah/SourceFlame/internal/errs"
	"github.com/Ankumeah/SourceFlame/internal/utils"
)

func r(a *app.App, root Route) {
	route := root.Sub("/r/{user}/{repo}")

	route.Handle(
		"GET /info/refs",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			service := r.URL.Query().Get("service")
			if service != "git-receive-pack" && service != "git-upload-pack" {
				utils.RespMsg(
					"Invalid service type",
					http.StatusBadRequest,
					w,
				)
				return
			}

			user := r.PathValue("user")
			if !utils.IsValidName(user) {
				utils.RespMsg(
					"Invalid username",
					http.StatusBadRequest,
					w,
				)
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

			w.Header().Add("Content-Type", "application/x-"+service+"-advertisement")

			if err := a.GitController.InfoRefs(
				ctx, filepath.Join(user, repo), service, w,
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
		},
	)

	route.Handle(
		"POST /git-upload-pack",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user := r.PathValue("user")
			if !utils.IsValidName(user) {
				utils.RespMsg(
					"Invalid username",
					http.StatusBadRequest,
					w,
				)
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

			if err := a.GitController.UploadPack(
				ctx, filepath.Join(user, repo), r.Body, w,
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
		},
	)

	route.Handle(
		"POST /git-receive-pack",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user := r.PathValue("user")
			if !utils.IsValidName(user) {
				utils.RespMsg(
					"Invalid username",
					http.StatusBadRequest,
					w,
				)
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

			if err := a.GitController.ReceivePack(
				ctx, filepath.Join(user, repo), r.Body, w,
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
		},
	)
}

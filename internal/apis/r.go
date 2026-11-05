package apis

import (
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/git"
	"github.com/Ankumeah/SourceFlame/internal/middlewares"
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

			if !handle(
				a.GitController.InfoRefs(
					ctx, filepath.Join(user, repo), service, w,
				), w,
			) {
				return
			}
		},
	)

	route.Handle(
		"GET /refs",
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

			refs, err := a.GitController.GetRefs(ctx, filepath.Join(user, repo))
			if !handle(err, w) {
				return
			}

			if !handle(
				json.NewEncoder(w).Encode(
					struct {
						Refs []git.Ref `json:"refs"`
					}{
						Refs: refs,
					},
				), w,
			) {
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

			if !handle(
				a.GitController.UploadPack(
					ctx, filepath.Join(user, repo), r.Body, w,
				), w,
			) {
				return
			}
		},
	)

	route.Handle(
		"GET /blob/{file...}",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			file := r.PathValue("file")
			if strings.TrimSpace(file) == "" {
				utils.RespMsg(
					"Empty file path",
					http.StatusBadRequest,
					w,
				)
				return
			}

			ref := r.URL.Query().Get("ref")
			if strings.TrimSpace(ref) == "" {
				utils.RespMsg(
					"Empty ref",
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

			content, err := a.GitController.GetBlob(ctx, filepath.Join(user, repo), ref, file)
			if !handle(err, w) {
				return
			}

			w.Write(content)
		},
	)

	route.Handle(
		"POST /git-receive-pack",
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			owner := r.PathValue("user")
			if !utils.IsValidName(owner) {
				utils.RespMsg(
					"Invalid username",
					http.StatusBadRequest,
					w,
				)
				return
			}

			user, ok := ctx.Value(middlewares.UserKey).(string)
			if !ok {
				utils.RespErr(errors.New("user set by middleware not string"), w)
				return
			}

			if user != owner {
				utils.RespMsg(
					"Must be owner to push",
					http.StatusForbidden,
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

			if !handle(
				a.GitController.ReceivePack(
					ctx, filepath.Join(owner, repo), r.Body, w,
				), w,
			) {
				return
			}
		}, func(h http.Handler) http.Handler { return middlewares.BasicAuth(h, a) },
	)
}

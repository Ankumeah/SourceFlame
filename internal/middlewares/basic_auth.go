package middlewares

import (
	"errors"
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	e "github.com/Ankumeah/SourceFlame/internal/errs"
	"github.com/Ankumeah/SourceFlame/internal/utils"
)

func BasicAuth(next http.Handler, app *app.App) http.Handler {
	return ParseBasicAuth(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()

				user, ok := ctx.Value(UserKey).(string)
				if !ok {
					utils.RespErr(errors.New("user set by middleware not string"), w)
					return
				}

				pass, ok := ctx.Value(PassKey).(string)
				if !ok {
					utils.RespErr(errors.New("pass set by middleware not string"), w)
					return
				}

				actualHash, err := app.DB.GetPassHash(ctx, user)
				if errors.Is(err, e.ErrInvalidUser) {
					utils.RespMsg(
						"Invalid user",
						http.StatusNotFound,
						w,
					)
					return
				} else if err != nil {
					utils.RespErr(err, w)
					return
				}

				valid, err := app.Hasher.CompareHash(actualHash, pass)
				if err != nil {
					utils.RespErr(errors.New("error while comparing hash"), w)
					return
				}

				if !valid {
					utils.RespMsg(
						"Invalid password",
						http.StatusForbidden,
						w,
					)
					return
				}

				next.ServeHTTP(w, r)
			},
		),
	)
}

package apis

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

func ping(a *app.App, mux *http.ServeMux) {
	mux.HandleFunc(
		"GET /ping",
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte("PONG")); err != nil {
				return
			}
		},
	)
}

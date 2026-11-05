package repo

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

const prefix = "/r"

func Handle(app *app.App, mux *http.ServeMux) {
	mux.HandleFunc(
		"GET "+prefix+"/{user}/{repo}/{$}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
		},
	)

	mux.HandleFunc(
		"GET "+prefix+"/{user}/{repo}/styles.css",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/css")
			w.Write(styles)
		},
	)

	mux.HandleFunc(
		"GET "+prefix+"/{user}/{repo}/script.js",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/javascript")
			w.Write(script)
		},
	)
}

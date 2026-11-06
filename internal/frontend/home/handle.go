package home

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

const prefix = ""

func Handle(app *app.App, mux *http.ServeMux) {
	mux.HandleFunc(
		"GET "+prefix+"/{$}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
		},
	)

	mux.HandleFunc(
		"GET "+prefix+"/styles.css",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/css")
			w.Write(styles)
		},
	)

	mux.HandleFunc(
		"GET "+prefix+"/script.js",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/javascript")
			w.Write(script)
		},
	)
}

package repo

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

const prefix = "/r/{user}/{repo}"

func Handle(app *app.App, mux *http.ServeMux) {
	for _, route := range []struct {
		path        string
		contentType string
		file        []byte
	}{
		{"/{$}", "text/html", index},
		{"/styles.css", "text/css", styles},
		{"/script.js", "text/javascript", script},
	} {
		mux.HandleFunc(
			"GET "+prefix+route.path,
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", route.contentType)
				w.Write(route.file)
			},
		)
	}
}

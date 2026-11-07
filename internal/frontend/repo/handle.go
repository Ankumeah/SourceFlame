package repo

import (
	"fmt"
	"net/http"
	"strings"

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

	for route, method := range map[string]string{
		"/info/refs":        "GET",
		"/git-receive-pack": "POST",
		"/git-upload-pack":  "POST",
	} {
		mux.HandleFunc(
			method+" "+prefix+route,
			func(w http.ResponseWriter, r *http.Request) {
				queries := r.URL.RawQuery
				if !(strings.TrimSpace(queries) == "") {
					queries = "?" + queries
				}
				fmt.Println(queries)

				http.Redirect(
					w, r,
					fmt.Sprintf(
						"/api/%s/r/%s/%s%s%s",
						app.Config.ApiVersion,
						r.PathValue("user"),
						r.PathValue("repo"),
						route,
						queries,
					),
					http.StatusTemporaryRedirect,
				)
			},
		)
	}
}

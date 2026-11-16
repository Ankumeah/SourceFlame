package apis

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

func Handle(a *app.App, mux *http.ServeMux) {
	apiRoute := NewRoute(mux, "/api")
	baseRoute := NewRoute(mux, "")

	ping(a, apiRoute)
	acc(a, apiRoute)
	repos(a, apiRoute)

	r(a, baseRoute)
}

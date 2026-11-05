package apis

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

func Handle(a *app.App, mux *http.ServeMux) {
	route := NewRoute(mux, "/api/"+a.Config.ApiVersion)

	ping(a, route)
	acc(a, route)
	repos(a, route)
	r(a, route)
}

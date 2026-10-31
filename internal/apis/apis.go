package apis

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/middlewares"
)

func Apis(a *app.App) http.Handler {
	mux := http.NewServeMux()
	route := NewRoute(mux, "/api/"+a.Config.ApiVersion, middlewares.Log)

	ping(a, route)
	acc(a, route)
	repos(a, route)
	r(a, route)

	return mux
}

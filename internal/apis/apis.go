package apis

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/middleware"
)

func Apis(a *app.App) http.Handler {
	mux := http.NewServeMux()

	ping(a, mux)

	return middleware.Log(mux)
}

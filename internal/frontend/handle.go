package frontend

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/frontend/home"
	"github.com/Ankumeah/SourceFlame/internal/frontend/repo"
)

func Handle(app *app.App, mux *http.ServeMux) {
	home.Handle(app, mux)
	repo.Handle(app, mux)
}

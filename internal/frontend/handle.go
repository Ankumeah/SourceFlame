package frontend

import (
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/frontend/repo"
)

func Handle(app *app.App, mux *http.ServeMux) {
	repo.Handle(app, mux)
}
